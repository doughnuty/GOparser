package parser

import (
	parser "github.com/doughnuty/GOparser"
	"testing"
)

func TestParse(t *testing.T) {
	config := parser.NewYaml()

	sources := []parser.Source{parser.NewEnvSource(parser.WithPrefix("CGO"), parser.WithStrippedPrefix("IDEA")),
		parser.NewYamlSource("file.yaml")}
	err := config.Load(sources...)
	if err != nil {
		t.Errorf("Bad parsing. Error message is: %v", err)
	}

}
