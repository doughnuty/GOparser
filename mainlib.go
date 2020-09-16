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
	Bool(def bool) bool
	Int(def int) int
	String(def string) string
	Float64(def float64) float64
	Duration(def time.Duration) time.Duration
	StringSlice(def []string) []string
	StringMap(def map[string]interface{}) map[string]interface{}
	Bytes() []byte
}

func CreateYamlParser() Object {
	var object Object
	object.Objects = make(map[string]interface{})
	return object
}

// Parse is a function used to fill an Object structure
func (object *Object) Parse(source string) error {
	// open
	file, err := os.Open(source)
	if err != nil {
		return err
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
	return object.createObject(text, 0)
}

// Get is a function to get value from the Object structure
func (object *Object) Get(path ...string) Subjects {
	var ans Answer
	// check if nil
	if object == nil {
		fmt.Println("Empty Object")
		return ans
	}

	// seek for the path recursively
	ans.val = ans.seekVal(*object, path)

	return ans
}
