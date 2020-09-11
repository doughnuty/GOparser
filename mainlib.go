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
	Bool(def bool) (bool, error)
	Int(def int) (int, error)
	String(def string) (string, error)
	Float64(def float64) (float64, error)
	Duration(def time.Duration) (time.Duration, error)
	StringSlice(def []string) ([]string, error)
	StringMap(def map[string]interface{}) (map[string]interface{}, error)
	Bytes() ([]byte, error)
}

func Init() Object {
	var object Object
	object.Objects = make(map[string]interface{})
	return object
}

// Parse is a function used to fill an Object structure
func (dest *Object) Parse(source string) error {
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
	return createObject(text, dest, 0)
}

// Get is a function to get value from the Object structure
func (dest *Object) Get(path ...string) Subjects {
	var ans Answer
	// check if nil
	if dest == nil {
		fmt.Println("Empty Object")
		return ans
	}

	// seek for the path recursively
	ans.val = ans.seekVal(*dest, path)

	return ans
}