package parser

import (
	"testing"
	"theRealParser/parser"
)

func TestParse(t *testing.T) {
	yaml := parser.NewYaml()

	err := yaml.Parse("file.yaml")
	if err != nil {
		t.Errorf("Bad parsing. Error message is: %v", err)
	}
}
