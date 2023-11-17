package docs

import (
	"flag"
	"github.com/grafana/agent/component/metadata"
	"strings"
	"testing"

	"github.com/grafana/agent/component"
	_ "github.com/grafana/agent/component/all"
	"github.com/grafana/agent/docs/generator"
	"github.com/stretchr/testify/require"
)

// Run the below generate command to automatically update the Markdown docs with generated content
//go:generate go test -fix-tests

var fixTestsFlag = flag.Bool("fix-tests", false, "update the test files with the current generated content")

func TestCompatibleComponentsSectionsUpdated(t *testing.T) {
	for _, name := range component.AllNames() {
		t.Run(name, func(t *testing.T) {
			generated, err := generator.GenerateCompatibleComponentsSection(name)
			require.NoError(t, err, "failed to generate references section for %q", name)

			if generated == "" {
				t.Skipf("no compatible components section defined for %q", name)
			}

			if *fixTestsFlag {
				err = generator.WriteCompatibleComponentsSection(name)
				require.NoError(t, err, "failed to write generated references section for %q", name)
				t.Log("updated the docs with generated content")
			}

			actual, err := generator.ReadCompatibleComponentsSection(name)
			require.NoError(t, err, "failed to read generated components section for %q, try running 'go generate ./docs'", name)
			require.Contains(
				t,
				actual,
				strings.TrimSpace(generated),
				"expected documentation for %q to contain generated references section, try running 'go generate ./docs'",
				name,
			)
		})
	}
}

func TestCompatibleComponentsPageUpdated(t *testing.T) {
	path := "sources/flow/reference/compatibility/_index.md"
	for _, typ := range metadata.AllTypes {
		t.Run(typ.Name, func(t *testing.T) {
			t.Run("exporters", func(t *testing.T) {
				runForGenerator(t, generator.NewExportersListGenerator(typ, path))
			})
			t.Run("consumers", func(t *testing.T) {
				runForGenerator(t, generator.NewConsumersListGenerator(typ, path))
			})
		})
	}
}

func runForGenerator(t *testing.T, g generator.DocsGenerator) {
	generated, err := g.Generate()
	require.NoError(t, err, "failed to generate: %q", g.Name())

	if generated == "" {
		t.Skipf("nothing generated for %q, skipping", g.Name())
	}

	if *fixTestsFlag {
		err = g.Write()
		require.NoError(t, err, "failed to write generated content for: %q", g.Name())
		t.Log("updated the docs with generated content", g.Name())
	}

	actual, err := g.Read()
	require.NoError(t, err, "failed to read existing generated docs for %q, try running 'go generate ./docs'", g.Name())
	require.Contains(
		t,
		actual,
		strings.TrimSpace(generated),
		"outdated docs detected when running %q, try updating with 'go generate ./docs'",
		g.Name(),
	)
}
