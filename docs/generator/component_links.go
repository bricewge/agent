package generator

import (
	"fmt"
	"github.com/grafana/agent/component/metadata"
)

type LinksToTypesGenerator struct {
	component string
}

func NewLinksToTypesGenerator(component string) *LinksToTypesGenerator {
	return &LinksToTypesGenerator{component: component}
}

func (l *LinksToTypesGenerator) Name() string {
	return fmt.Sprintf("generator of links to types for %q reference page", l.component)
}

func (l *LinksToTypesGenerator) Generate() (string, error) {
	meta, err := metadata.ForComponent(l.component)
	if err != nil {
		return "", err
	}
	if meta.Empty() {
		return "", nil
	}

	heading := "\n## Compatible components\n\n"
	acceptingSection := acceptingComponentsSection(l.component, meta)
	outputSection := outputComponentsSection(l.component, meta)

	if acceptingSection == "" && outputSection == "" {
		return "", nil
	}

	note := "\nNote that connecting some components may not be feasible or components may require further " +
		"configuration to make the connection work correctly. " +
		"Please refer to the linked documentation for more details.\n\n"

	return heading + acceptingSection + outputSection + note, nil
}

func (l *LinksToTypesGenerator) Read() (string, error) {
	content, err := readBetweenMarkers(l.startMarker(), l.endMarker(), l.pathToComponentMarkdown())
	if err != nil {
		return "", fmt.Errorf("failed to read existing content for %q: %w", l.Name(), err)
	}
	return content, err
}

func (l *LinksToTypesGenerator) Write() error {
	newSection, err := l.Generate()
	if err != nil {
		return err
	}
	newSection = "\n" + newSection + "\n"
	return writeBetweenMarkers(l.startMarker(), l.endMarker(), l.pathToComponentMarkdown(), newSection, true)
}

func (l *LinksToTypesGenerator) startMarker() string {
	return "<!-- START GENERATED COMPATIBLE COMPONENTS -->"
}

func (l *LinksToTypesGenerator) endMarker() string {
	return "<!-- END GENERATED COMPATIBLE COMPONENTS -->"
}

func (l *LinksToTypesGenerator) pathToComponentMarkdown() string {
	return fmt.Sprintf("sources/flow/reference/components/%s.md", l.component)
}

func outputComponentsSection(name string, meta metadata.Metadata) string {
	section := ""
	for _, outputDataType := range meta.Exports {
		if list := allComponentsThatAccept(outputDataType); len(list) > 0 {
			section += fmt.Sprintf("- Components that accept [%s]({{< relref \"../compatibility\" >}})\n", outputDataType.Name)
		}
	}
	if section != "" {
		section = fmt.Sprintf("`%s` exports can be consumed by the following components:\n\n", name) + section
	}
	return section
}

func acceptingComponentsSection(componentName string, meta metadata.Metadata) string {
	section := ""
	for _, acceptedDataType := range meta.Accepts {
		if list := allComponentsThatExport(acceptedDataType); len(list) > 0 {
			section += fmt.Sprintf("- Components that export [%s]({{< relref \"../compatibility\" >}})\n", acceptedDataType.Name)
		}
	}
	if section != "" {
		section = fmt.Sprintf("`%s` can accept arguments from the following components:\n\n", componentName) + section + "\n"
	}
	return section
}
