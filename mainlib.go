package parser

import (
	"bufio"
	"time"
	//	"bytes"
	//	"errors"
	"fmt"
	// "io"
	"os"
)

// Object is  a type to hold values
type Object struct {
	Objects map[string]interface{}
}

// Answer is a structure
type Answer struct {
	val interface{}
}

// Subjects is an interface with supported type of values
type Subjects interface {
	Bool() bool
	Int() int
	String() string
	Float64() float64
	Duration() time.Duration
	StringSlice() []string
	StringMap() map[string]interface{}
	Bytes() []byte
}

// Parse is a function used to fill an Object structure
func Parse(source string, dest *Object) {
	// open
	file, err := os.Open(source)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// save values
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// createObject recursively to fill dest
	createObject(text, dest, 0)

	return
}

// Get is a function to get value from the Object structure
func Get(source *Object, path ...string) Subjects {
	var ans Answer
	// check if nil
	if source == nil {
		fmt.Println("Empty Object")
		return ans
	}

	// seek for the path recursively
	ans.val = ans.seekVal(*source, path)

	return ans
}