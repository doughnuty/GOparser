package main

import (
	"fmt"
	"theRealParser/parser"
)

func main() {
	yaml := parser.NewYaml()

	err := yaml.Parse("file.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(yaml.Get("student", "clubs", "Art", "ToDo").StringSlice(nil))
	}
}
