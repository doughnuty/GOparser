package parser

import (
	"bufio"
	//	"bytes"
	//	"errors"
	"fmt"
	// "io"
	"os"
	"reflect"
	// "unicode"
	// "parser.go"
	"strconv"
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
	//Float64(def float64) float64
	// Duration(def time.Duration) time.Duration
	StringSlice() []string
	StringMap() map[string]interface{}
	//Bytes() []byte
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
	ans.val = ans.seekVal(source, path)

	return ans
}

// String gets a string type
func (ans Answer) String() (ret string) {
	// if empty panic
	ret = fmt.Sprintf("%v", ans.val)
	if ret == "<nil>" {
		panic("No match")
	}

	// if not a string notify
	if reflect.TypeOf(ans.val) != reflect.TypeOf(ret) {
		fmt.Println("Mismatched types", reflect.TypeOf(ans.val), "and", reflect.TypeOf(ret))
	}
	return ret
}

// Int gets an integer type
func (ans Answer) Int() int {
	// if empty panic
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		panic("No match")
	}

	// else convert (if incompatible strconv will panic)
	ret, err := strconv.Atoi(temp)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return ret
}

// Bool gets a boolean type
func (ans Answer) Bool() bool {
	// if empty panic
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		panic("No match")
	}

	// else convert (if incompatible strconv will panic)
	ret, err := strconv.ParseBool(temp)
	if err != nil {
		panic(err)
	}
	return ret
}

func (ans Answer) StringMap() map[string]interface{} {
	ret := new(map[string]interface{})

	// if empty panic
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		panic("No match")
	}

	// if not an Object return
	var tempObject Object

	if reflect.TypeOf(ans.val) != reflect.TypeOf(tempObject) {
		fmt.Println("Mismatched types", reflect.TypeOf(ans.val), "and", reflect.TypeOf(tempObject))
		return *ret
	}

	// if Object save Object.Objects value
	tempObject = ans.val.(Object)
	*ret = tempObject.Objects

	return *ret
}

func (ans Answer) StringSlice() []string {
	var ret []string
	interf := new([]interface{})

	// if empty panic
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		panic("No match")
	}

	// if not a slice return
	if reflect.TypeOf(ans.val) != reflect.TypeOf(*interf) {
		fmt.Println("Mismatched types", reflect.TypeOf(ans.val), "and", reflect.TypeOf(*interf))
		return ret
	}
	for i, val := range ans.val.([]interface{}) {
		// if not a string return
		if reflect.TypeOf(ans.val.([]interface{})[i]) != reflect.TypeOf(temp) {
			fmt.Println("Mismatched types", reflect.TypeOf(ans.val.([]interface{})[i]), "and", reflect.TypeOf(temp))
			return ret
		}
		// if string append
		ret = append(ret, val.(string))
	}

	return ret
}
