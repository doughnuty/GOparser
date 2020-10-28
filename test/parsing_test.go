package parser

import (
	parser "github.com/doughnuty/GOparser"
	"testing"
)

func TestParse(t *testing.T) {
	config := parser.NewYaml()

	yaml := parser.NewYamlSource("file.yaml")
	env := parser.NewEnvSource(parser.WithPrefix("CGO"))

	err := config.Load(yaml, env)
	if err != nil {
		t.Errorf("Bad parsing. Error message is: %v", err)
	}

}
