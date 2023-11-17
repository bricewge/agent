package metadata

import (
	"fmt"
	"reflect"

	"github.com/grafana/agent/component"
	_ "github.com/grafana/agent/component/all"
	"github.com/grafana/agent/component/common/loki"
	"github.com/grafana/agent/component/discovery"
)

//TODO(thampiotr): Instead of metadata package reaching into registry, we'll migrate to using a YAML metadata file that
//				   contains information about all the components. This file will be generated separately from the
//				   registry and can be used by other tools.

type Type struct {
	Name string
	// Returns true if provided args include this type (including nested structs)
	ExistsInArgsFn func(args component.Arguments) bool
	// Returns true if provided exports include this type (including nested structs)
	ExistsInExportsFn func(exports component.Exports) bool
}

var (
	// TypeTargets represents things that need to be scraped.
	TypeTargets = Type{
		Name: "Targets",
		ExistsInArgsFn: func(args component.Arguments) bool {
			return hasFieldOfType(args, reflect.TypeOf([]discovery.Target{}))
		},
		ExistsInExportsFn: func(exports component.Exports) bool {
			return hasFieldOfType(exports, reflect.TypeOf([]discovery.Target{}))
		},
	}

	// TypeLokiLogs represent logs in Loki format
	TypeLokiLogs = Type{
		Name: "Loki `LogsReceiver`",
		ExistsInArgsFn: func(args component.Arguments) bool {
			return hasFieldOfType(args, reflect.TypeOf([]loki.LogsReceiver{}))
		},
		ExistsInExportsFn: func(exports component.Exports) bool {
			return hasFieldOfType(exports, reflect.TypeOf(loki.NewLogsReceiver()))
		},
	}

	//TODO(thampiotr): add more types
	//DataTypeOTELTelemetry     = Type("OTEL Telemetry")
	//DataTypePromMetrics       = Type("Prometheus Metrics")
	//DataTypePyroscopeProfiles = Type("Pyroscope Profiles")

	AllTypes = []Type{
		TypeTargets,
		TypeLokiLogs,
	}
)

type Metadata struct {
	Accepts []Type
	Exports []Type
}

func (m Metadata) Empty() bool {
	return len(m.Accepts) == 0 && len(m.Exports) == 0
}

func (m Metadata) AcceptsType(t Type) bool {
	for _, a := range m.Accepts {
		if a.Name == t.Name {
			return true
		}
	}
	return false
}

func (m Metadata) ExportsType(t Type) bool {
	for _, o := range m.Exports {
		if o.Name == t.Name {
			return true
		}
	}
	return false
}

func ForComponent(name string) (Metadata, error) {
	reg, ok := component.Get(name)
	if !ok {
		return Metadata{}, fmt.Errorf("could not find component %q", name)
	}
	return inferMetadata(reg.Args, reg.Exports), nil
}

func inferMetadata(args component.Arguments, exports component.Exports) Metadata {
	m := Metadata{}
	for _, t := range AllTypes {
		if t.ExistsInArgsFn(args) {
			m.Accepts = append(m.Accepts, t)
		}
		if t.ExistsInExportsFn(exports) {
			m.Exports = append(m.Exports, t)
		}
	}
	return m
}

func hasFieldOfType(obj interface{}, fieldType reflect.Type) bool {
	objValue := reflect.ValueOf(obj)

	// If the object is a pointer, dereference it
	for objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	// If the object is not a struct or interface, return false
	if objValue.Kind() != reflect.Struct && objValue.Kind() != reflect.Interface {
		return false
	}

	for i := 0; i < objValue.NumField(); i++ {
		fv := objValue.Field(i)
		ft := fv.Type()

		// If the field type matches the given type, return true
		if ft == fieldType {
			return true
		}

		if fv.Kind() == reflect.Interface && fieldType.AssignableTo(ft) {
			return true
		}

		// If the field is a struct, recursively check its fields
		if fv.Kind() == reflect.Struct {
			if hasFieldOfType(fv.Interface(), fieldType) {
				return true
			}
		}
	}

	return false
}
