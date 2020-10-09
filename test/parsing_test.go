package parser

import (
	"GOparser"
	"testing"
)

func TestParse(t *testing.T) {
	yaml := GOparser.NewYaml()

	err := yaml.Parse("file.yaml")
	if err != nil {
		t.Errorf("Bad parsing. Error message is: %v", err)
	}
}
