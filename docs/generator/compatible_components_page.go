package generator

import (
	"fmt"
	"strings"

	"github.com/grafana/agent/component/metadata"
)

type CompatibleComponentsListGenerator struct {
	filePath    string
	t           metadata.Type
	sectionName string
	generateFn  func() string
}

func NewExportersListGenerator(t metadata.Type, filePath string) *CompatibleComponentsListGenerator {
	return &CompatibleComponentsListGenerator{
		filePath:    filePath,
		t:           t,
		sectionName: "exporters",
		generateFn:  func() string { return listOfComponentsExporting(t) },
	}
}

func NewConsumersListGenerator(t metadata.Type, filePath string) *CompatibleComponentsListGenerator {
	return &CompatibleComponentsListGenerator{
		filePath:    filePath,
		t:           t,
		sectionName: "consumers",
		generateFn:  func() string { return listOfComponentsAccepting(t) },
	}
}

func (c *CompatibleComponentsListGenerator) Name() string {
	return fmt.Sprintf("generator of %s section for %q in %q", c.sectionName, c.t.Name, c.filePath)
}

func (c *CompatibleComponentsListGenerator) Generate() (string, error) {
	return c.generateFn(), nil
}

func (c *CompatibleComponentsListGenerator) Read() (string, error) {
	content, err := readBetweenMarkers(c.startMarker(), c.endMarker(), c.filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read existing content for %q: %w", c.Name(), err)
	}
	return content, err
}

func (c *CompatibleComponentsListGenerator) Write() error {
	newSection, err := c.Generate()
	if err != nil {
		return err
	}
	newSection = "\n" + newSection + "\n"
	return writeBetweenMarkers(c.startMarker(), c.endMarker(), c.filePath, newSection, false)
}

func (c *CompatibleComponentsListGenerator) startMarker() string {
	return fmt.Sprintf("<!-- START GENERATED SECTION: %s OF %s -->", strings.ToUpper(c.sectionName), c.t.Name)
}

func (c *CompatibleComponentsListGenerator) endMarker() string {
	return fmt.Sprintf("<!-- END GENERATED SECTION: %s OF %s -->", strings.ToUpper(c.sectionName), c.t.Name)
}

func listOfComponentsAccepting(dataType metadata.Type) string {
	return listOfLinksToComponents(allComponentsThatAccept(dataType))
}

func listOfComponentsExporting(dataType metadata.Type) string {
	return listOfLinksToComponents(allComponentsThatExport(dataType))
}

func listOfLinksToComponents(components []string) string {
	str := ""
	groups := make(map[string][]string)

	for _, component := range components {
		parts := strings.SplitN(component, ".", 2)
		namespace := parts[0]
		groups[namespace] = append(groups[namespace], component)
	}

	for namespace, components := range groups {
		str += fmt.Sprintf("\n{{< collapse title=%q >}}\n", namespace)
		for _, component := range components {
			str += fmt.Sprintf("- [%[1]s]({{< relref \"../components/%[1]s.md\" >}})\n", component)
		}
		str += fmt.Sprintf("{{< /collapse >}}\n")
	}
	return str
}
