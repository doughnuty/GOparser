# Go yaml files parser
Small parser for yaml files written in Go. 

## Synopsis 
```GO
package main

import (
	"fmt"
	"theRealParser/parser"
)

func main() {
	yaml := parser.NewYaml()
	err := yaml.Parse("file.yaml")
	if err != nil {
		fmt.Println(err)
	}

	ans := yaml.Get("student", "personal", "Name").String("hewwo")
	fmt.Println(ans)
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
