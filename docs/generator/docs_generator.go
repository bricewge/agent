package generator

import (
	"github.com/grafana/agent/component"
	"github.com/grafana/agent/component/metadata"
)

type DocsGenerator interface {
	Name() string
	Generate() (string, error)
	Read() (string, error)
	Write() error
}

func allComponentsThat(f func(meta metadata.Metadata) bool) []string {
	var result []string
	for _, name := range component.AllNames() {
		meta, err := metadata.ForComponent(name)
		if err != nil {
			panic(err) // should never happen
		}

		if f(meta) {
			result = append(result, name)
		}
	}
	return result
}

func allComponentsThatExport(dataType metadata.Type) []string {
	return allComponentsThat(func(meta metadata.Metadata) bool {
		return meta.ExportsType(dataType)
	})
}

func allComponentsThatAccept(dataType metadata.Type) []string {
	return allComponentsThat(func(meta metadata.Metadata) bool {
		return meta.AcceptsType(dataType)
	})
}
