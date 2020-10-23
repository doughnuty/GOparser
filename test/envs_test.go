package parser

import (
	"github.com/doughnuty/GOparser"
	"testing"
)

func TestEnv(t *testing.T) {
	yamlEnv := GOparser.NewYaml()

	err := yamlEnv.ParseDotEnv(GOparser.WithStrippedPrefix("IDEA"), GOparser.WithPrefix("CGO"))
	if err != nil {
		t.Errorf("Bad parsing. Error message is: %v", err)
	}

	testMap := yamlEnv.Get("cgo").StringMap(nil)
	if testMap == nil {
		t.Errorf("unsuccessfull parse")
		t.Errorf("Parsed as %v", yamlEnv.Get("cgo"))
	}

	testString := yamlEnv.Get("initial", "directory").String("")
	if testString == "" {
		t.Errorf("Error parsing a string")
	}
}
