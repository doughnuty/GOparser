package parser

import (
	parser "github.com/doughnuty/GOparser"
	"testing"
)

func TestEnv(t *testing.T) {
	config := parser.NewYaml()

	env := parser.NewEnvSource(parser.WithPrefix("CGO"), parser.WithStrippedPrefix("IDEA"))

	err := config.Load(env)
	if err != nil {
		t.Errorf("Bad parsing. Error message is: %v", err)
	}

	testMap := config.Get("cgo").StringMap(nil)
	if testMap == nil {
		t.Errorf("unsuccessfull parse")
		t.Errorf("Parsed as %v", config.Get("cgo"))
	}

	testString := config.Get("initial", "directory").String("")
	if testString == "" {
		t.Errorf("Error parsing a string")
	}
}
