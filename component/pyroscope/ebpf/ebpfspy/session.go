//go:build linux

// Package ebpfspy provides integration with Linux eBPF. It is a rough copy of profile.py from BCC tools:
//
//	https://github.com/iovisor/bcc/blob/master/tools/profile.py
package ebpfspy

import (
	_ "embed"
	"encoding/binary"
	"fmt"
	"reflect"
	"sync"
	"unsafe"

	"github.com/cilium/ebpf"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grafana/agent/component/pyroscope/ebpf/ebpfspy/sd"
	"github.com/pyroscope-io/pyroscope/pkg/agent/ebpfspy/cpuonline"
	"golang.org/x/sys/unix"
)

//go:generate make -C bpf get-headers
//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang -cflags "-O2 -Wall -fpie -Wno-unused-variable -Wno-unused-function" profile bpf/profile.bpf.c -- -I./bpf/libbpf -I./bpf/vmlinux/

type Session struct {
	logger           log.Logger
	pid              int
	sampleRate       uint32
	pidCacheSize     int
	elfCacheSize     int
	serviceDiscovery *sd.TargetFinder

	perfEvents []*perfEvent

	symCache *symbolCache

	bpf profileObjects

	modMutex sync.Mutex

	roundNumber int
}

func NewSession(
	logger log.Logger,
	serviceDiscovery *sd.TargetFinder,
	sampleRate uint32,
	pidCacheSize int,
	elfCacheSize int,
) (*Session, error) {

	symCache, err := newSymbolCache(logger, pidCacheSize, elfCacheSize)
	if err != nil {
		return nil, err
	}

	return &Session{
		logger:           logger,
		pid:              -1,
		symCache:         symCache,
		sampleRate:       sampleRate,
		pidCacheSize:     pidCacheSize,
		elfCacheSize:     elfCacheSize,
		serviceDiscovery: serviceDiscovery,
	}, nil
}

func (s *Session) Start() error {
	var err error
	if err = unix.Setrlimit(unix.RLIMIT_MEMLOCK, &unix.Rlimit{
		Cur: unix.RLIM_INFINITY,
		Max: unix.RLIM_INFINITY,
	}); err != nil {
		return err
	}

	s.modMutex.Lock()
	defer s.modMutex.Unlock()

	opts := &ebpf.CollectionOptions{}
	if err := loadProfileObjects(&s.bpf, opts); err != nil {
		return fmt.Errorf("load bpf objects: %w", err)
	}
	if err = s.initArgs(); err != nil {
		return fmt.Errorf("init bpf args: %w", err)
	}
	if err = s.attachPerfEvents(); err != nil {
		return fmt.Errorf("attach perf events: %w", err)
	}
	return nil
}

func (s *Session) Reset(cb func(t *sd.Target, stack []string, value uint64, pid uint32) error) error {
	level.Debug(s.logger).Log("msg", "ebpf session reset")
	s.modMutex.Lock()
	defer s.modMutex.Unlock()

	s.roundNumber += 1

	keys, values, batch, err := s.getCountsMapValues()
	if err != nil {
		return fmt.Errorf("get counts map: %w", err)
	}

	type sf struct {
		pid    uint32
		count  uint32
		kStack []byte
		uStack []byte
		comm   string
		labels *sd.Target
	}
	var sfs []sf
	knownStacks := map[uint32]bool{}
	for i := range keys {
		ck := &keys[i]
		value := values[i]

		if ck.UserStack >= 0 {
			knownStacks[uint32(ck.UserStack)] = true
		}
		if ck.KernStack >= 0 {
			knownStacks[uint32(ck.KernStack)] = true
		}
		labels := s.serviceDiscovery.FindTarget(ck.Pid)
		if labels == nil {
			continue
		}
		uStack := s.getStack(ck.UserStack)
		kStack := s.getStack(ck.KernStack)
		sfs = append(sfs, sf{
			pid:    ck.Pid,
			uStack: uStack,
			kStack: kStack,
			count:  value,
			comm:   getComm(ck),
			labels: labels,
		})
	}

	sb := stackBuilder{}
	for _, it := range sfs {
		sb.rest()
		sb.append(it.comm)
		s.walkStack(&sb, it.uStack, it.pid)
		s.walkStack(&sb, it.kStack, 0)
		reverse(sb.stack)
		err = cb(it.labels, sb.stack, uint64(it.count), it.pid)
		if err != nil {
			return err
		}
	}
	if err = s.clearCountsMap(keys, batch); err != nil {
		return fmt.Errorf("clear counts map %w", err)
	}
	if err = s.clearStacksMap(knownStacks); err != nil {
		return fmt.Errorf("clear stacks map %w", err)
	}
	return nil
}

func (s *Session) Stop() {
	for _, pe := range s.perfEvents {
		_ = pe.Close()
	}
	s.bpf.Close()
}

func (s *Session) initArgs() error {
	var zero uint32
	var tgidFilter uint32
	if s.pid <= 0 {
		tgidFilter = 0
	} else {
		tgidFilter = uint32(s.pid)
	}
	arg := &profileBssArg{
		TgidFilter: tgidFilter,
	}
	if err := s.bpf.Args.Update(&zero, arg, 0); err != nil {
		return fmt.Errorf("init args fail: %w", err)
	}
	return nil
}

func (s *Session) attachPerfEvents() error {
	var cpus []uint
	var err error
	if cpus, err = cpuonline.Get(); err != nil {
		return fmt.Errorf("get cpuonline: %w", err)
	}
	for _, cpu := range cpus {
		pe, err := newPerfEvent(int(cpu), int(s.sampleRate))
		if err != nil {
			return fmt.Errorf("new perf event: %w", err)
		}
		s.perfEvents = append(s.perfEvents, pe)

		err = pe.attachPerfEvent(s.bpf.profilePrograms.DoPerfEvent)
		if err != nil {
			return fmt.Errorf("attach perf event: %w", err)
		}
	}
	return nil
}

func (s *Session) getStack(stackId int64) []byte {
	if stackId < 0 {
		return nil
	}
	stackIdU32 := uint32(stackId)
	res, err := s.bpf.profileMaps.Stacks.LookupBytes(stackIdU32)
	if err != nil {
		return nil
	}
	return res
}

func (s *Session) walkStack(sb *stackBuilder, stack []byte, pid uint32) {
	if len(stack) == 0 {
		return
	}
	var stackFrames []string
	for i := 0; i < 127; i++ {
		it := stack[i*8 : i*8+8]
		ip := binary.LittleEndian.Uint64(it)
		if ip == 0 {
			break
		}
		sym := s.symCache.resolve(pid, ip, s.roundNumber)
		name := "[unknown]"
		if sym != nil && sym.Name != "" {
			name = sym.Name
		}
		stackFrames = append(stackFrames, name)
	}
	reverse(stackFrames)
	for _, s := range stackFrames {
		sb.append(s)
	}
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func getComm(k *profileSampleKey) string {
	res := ""
	sh := (*reflect.StringHeader)(unsafe.Pointer(&res))
	sh.Data = uintptr(unsafe.Pointer(&k.Comm[0]))
	for _, c := range k.Comm {
		if c != 0 {
			sh.Len++
		} else {
			break
		}
	}
	return res
}

type stackBuilder struct {
	stack []string
}

func (s *stackBuilder) rest() {
	s.stack = s.stack[:0]
}

func (s *stackBuilder) append(sym string) {
	s.stack = append(s.stack, sym)
}