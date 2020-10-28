# Go config parser 
Small parser for yaml files and environmental vars written in Go. 

## Synopsis 
```GO
package main

import (
	"fmt"
	parser "github.com/doughnuty/GOparser"
)

func main() {
	config := parser.NewYaml()

	yaml := parser.NewYamlSource("file.yaml")
	env := parser.NewEnvSource(parser.WithPrefix("CGO"))

	err := config.Load(yaml, env)
	if err != nil {
		...
	}

    str := config.Get("server").StringMap(nil)
}
```

## Overview
Function Parse writes contents of the file to structure Object. 
Then, the structure is passed to the function Get, which returns the interface. The type of the value is specified by one of the methods mentioned beneath:

Currently supports such methods as 	
 * Bool(def bool) (bool, error)
 * Int(def int) (int, error)
 * String(def string) (string, error)
 * Float64(def float64) (float64, error)
 * Duration(def time.Duration) (time.Duration, error)
 * StringSlice(def []string) ([]string, error)
 * StringMap(def map[string]interface{}) (map[string]interface{}, error)
 * Bytes() ([]byte, error)