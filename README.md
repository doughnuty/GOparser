# Go yaml files parser
Small parser for yaml files written in Go. 

## Synopsis 
```GO
package main

import (
	"fmt"
	"parser"
)

func main() {
	var yaml parser.Object
	yaml.Objects = make(map[string]interface{})
	parser.Parse("PathToFile", &yaml)
	//fmt.Println(yaml)
	ans := parser.Get(&yaml, "value1", "value2", "etc").String()
	fmt.Println(ans)
}
```

## Overview
Function Parse writes contents of the file to structure Object. 
Then, the structure is passed to the function Get, which returns the interface. The type of the value is specified by one of the methods mentioned beneath
Currently supports such methods as 	
  * Bool() bool
  * Int() int
  * String() string
  * StringSlice() []string
  * StringMap() map[string]interface{}
