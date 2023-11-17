package generator

import (
	"bytes"
	"fmt"
	"github.com/grafana/agent/component/metadata"
	"os"
)

type CompatibleComponentsListGenerator struct {
	filePath  string
	t         metadata.Type
	exporters bool
}

func NewExportersListGenerator(t metadata.Type, filePath string) *CompatibleComponentsListGenerator {
	return &CompatibleComponentsListGenerator{
		filePath:  filePath,
		t:         t,
		exporters: true,
	}
}

func NewConsumersListGenerator(t metadata.Type, filePath string) *CompatibleComponentsListGenerator {
	return &CompatibleComponentsListGenerator{
		filePath:  filePath,
		t:         t,
		exporters: false,
	}
}

func (c *CompatibleComponentsListGenerator) Name() string {
	expImp := "importers"
	if c.exporters {
		expImp = "exporters"
	}
	return fmt.Sprintf("generator of %s section for %q in %q", expImp, c.t.Name, c.filePath)
}

func (c *CompatibleComponentsListGenerator) Generate() (string, error) {
	return "dummy\ndummy\ndummy\ndummy\ndummy\n", nil
}

func (c *CompatibleComponentsListGenerator) Read() (string, error) {
	fileContents, err := os.ReadFile(c.filePath)
	if err != nil {
		return "", err
	}

	startMarker := c.startMarkerBytes()
	endMarker := c.endMarkerBytes()
	startIndex := bytes.Index(fileContents, startMarker)
	endIndex := bytes.Index(fileContents, endMarker)
	if startIndex == -1 || endIndex == -1 {
		return "", fmt.Errorf("existing section not found by %q", c.Name())
	}

	return string(fileContents[startIndex+len(startMarker) : endIndex]), nil

}

func (c *CompatibleComponentsListGenerator) Write() error {
	newSection, err := c.Generate()
	if err != nil {
		return err
	}
	fileContents, err := os.ReadFile(c.filePath)
	if err != nil {
		return err
	}

	startMarker := c.startMarkerBytes()
	endMarker := c.endMarkerBytes()
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

	err = os.WriteFile(c.filePath, newFileContents, 0644)
	return err
}

func (c *CompatibleComponentsListGenerator) startMarkerBytes() []byte {
	return []byte(fmt.Sprintf("<!-- START GENERATED SECTION: EXPORTERS OF %s -->", c.t.Name))
}

func (c *CompatibleComponentsListGenerator) endMarkerBytes() []byte {
	return []byte(fmt.Sprintf("<!-- END GENERATED SECTION: EXPORTERS OF %s -->", c.t.Name))
}
