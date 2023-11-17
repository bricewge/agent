package generator

import (
	"bytes"
	"fmt"
	"github.com/grafana/agent/component/metadata"
	"os"
)

const (
	startDelimiter = "<!-- START GENERATED COMPATIBLE COMPONENTS -->"
	endDelimiter   = "<!-- END GENERATED COMPATIBLE COMPONENTS -->"
)

func WriteCompatibleComponentsSection(componentName string) error {
	filePath := pathToComponentMarkdown(componentName)
	newSection, err := GenerateCompatibleComponentsSection(componentName)
	if err != nil {
		return err
	}
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	startMarker := startMarkerBytes()
	endMarker := endMarkerBytes()
	replacement := append(append(startMarker, []byte(newSection)...), endMarker...)

	startIndex := bytes.Index(fileContents, startMarker)
	endIndex := bytes.Index(fileContents, endMarker)
	var newFileContents []byte
	if startIndex == -1 || endIndex == -1 {
		// Append the new section to the end of the file
		newFileContents = append(fileContents, append([]byte("\n"), replacement...)...)
	} else {
		// Replace the section with the new content
		newFileContents = append(fileContents[:startIndex], replacement...)
		newFileContents = append(newFileContents, fileContents[endIndex+len(endMarker):]...)
	}

	err = os.WriteFile(filePath, newFileContents, 0644)
	return err
}

func ReadCompatibleComponentsSection(componentName string) (string, error) {
	filePath := pathToComponentMarkdown(componentName)
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	startMarker := startMarkerBytes()
	endMarker := endMarkerBytes()
	startIndex := bytes.Index(fileContents, startMarker)
	endIndex := bytes.Index(fileContents, endMarker)
	if startIndex == -1 || endIndex == -1 {
		return "", fmt.Errorf("compatible components section not found in %q", filePath)
	}

	return string(fileContents[startIndex+len(startMarker) : endIndex]), nil
}

func GenerateCompatibleComponentsSection(componentName string) (string, error) {
	meta, err := metadata.ForComponent(componentName)
	if err != nil {
		return "", err
	}
	if meta.Empty() {
		return "", nil
	}

	heading := "\n## Compatible components\n\n"
	acceptingSection := acceptingComponentsSection(componentName, meta)
	outputSection := outputComponentsSection(componentName, meta)

	if acceptingSection == "" && outputSection == "" {
		return "", nil
	}

	note := "\nNote that connecting some components may not be feasible or components may require further " +
		"configuration to make the connection work correctly. " +
		"Please refer to the linked documentation for more details.\n\n"

	return heading + acceptingSection + outputSection + note, nil
}

func outputComponentsSection(name string, meta metadata.Metadata) string {
	section := ""
	for _, outputDataType := range meta.Exports {
		if list := listOfComponentsAccepting(outputDataType); list != "" {
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
		if list := listOfComponentsExporting(acceptedDataType); list != "" {
			section += fmt.Sprintf("- Components that export [%s]({{< relref \"../compatibility\" >}})\n", acceptedDataType.Name)
		}
	}
	if section != "" {
		section = fmt.Sprintf("`%s` can accept arguments from the following components:\n\n", componentName) + section + "\n"
	}
	return section
}

func listOfComponentsAccepting(dataType metadata.Type) string {
	return listOfLinksToComponents(allComponentsThatAccept(dataType))
}

func listOfComponentsExporting(dataType metadata.Type) string {
	return listOfLinksToComponents(allComponentsThatExport(dataType))
}

func listOfLinksToComponents(components []string) string {
	str := ""
	for _, comp := range components {
		str += fmt.Sprintf("  - [`%[1]s`]({{< relref \"../components/%[1]s.md\" >}})\n", comp)
	}
	return str
}

func pathToComponentMarkdown(name string) string {
	return fmt.Sprintf("sources/flow/reference/components/%s.md", name)
}

func endMarkerBytes() []byte {
	return []byte(endDelimiter + "\n")
}

func startMarkerBytes() []byte {
	return []byte(startDelimiter + "\n")
}
