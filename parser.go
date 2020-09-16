package parser

import (
	"errors"
	"strings"
)

// Checks if proper spacing was applied
func spacing(line string, space int) int {
	for i, char := range line {
		if char != ' ' && i < space { // if not enough spacing exit
			/*fmt.Printf("Spacing Error on line %s\n", line)*/
			return 0
		} else if char == ' ' && i > space { // if too much spacing skip
			/*if i%2 > 0 {
				fmt.Printf("At line \"%s\"\n", line)
				panic("Spacing Error")
			}*/
			return 1
		} else if char != ' ' { // if came the contents break
			return 2
		}
	}
	return 2
}

// Checks if buf has yaml-like arrays.
// If so, returns the array and "true"
// If not, returns "false" and empty
func arrayCheck(buf []string, space int) (bool, []interface{}) {
	array := new([]interface{})
	for _, line := range buf {
		// check spacing (should make function?)
		spacing := spacing(line, space)
		if spacing == 0 || spacing == 1 { // if not enough return
			if (*array) == nil {
				return false, *array
			}
			return true, *array
		}

		// check for '-' character with split
		split := strings.SplitN(line, "- ", 2)

		// if len == 1 return
		if len(split) <= 1 {
			if (*array) == nil {
				return false, *array
			}
			return true, *array
		}
		// check if it's an array of Objects
		objects := strings.SplitN(split[1], ":", 2)
		if len(objects) == 2 {
			//if it is, create an Object and append it to the array
			newObject := CreateYamlParser()
			err := newObject.createObject(buf, space)
			if err != nil {
				return false, *array
			}
			*array = append(*array, newObject)
		} else {
			//if it's not append a string
			*array = append(*array, split[1])
		}
	}
	return true, *array
}

// Fills Object struct with values taken from buf.
// Uses yaml syntax, tracks spacing with space integer
func (object Object) createObject(buf []string, space int) error {
	// foreach line
	for lineNum, line := range buf {

		// check if spacing applied
		spacing := spacing(line, space)
		if spacing == 0 { // if not enough return
			return nil
		} else if spacing == 1 { // if too much spacing skip
			continue
		}

		split := strings.SplitN(line, ": ", 2) // search for ':' character
		// if not found return
		if len(split) <= 1 {
			split = strings.SplitN(line, ":", 2)
			if len(split) <= 1 {
				return errors.New("formatting error")
			}
		}

		// if key followed with word add it as value and search for next key - ex. foo: bar
		if len(split[1]) > 0 {
			// fmt.Printf("Assigning %s to %s\n", split[1], split[0][space:])
			object.Objects[split[0][space:]] = split[1]
		} else {
			lineNum++ // increase line counter

			if len(buf) <= lineNum { // check if line exist
				return errors.New("not enough values in the file")
			}

			// check if next line starts with '-' to make an array
			isArray, array := arrayCheck(buf[lineNum:], space+2)

			key := split[0][space:]
			arrayObjectSplit := strings.SplitN(split[0], "- ", 2)
			if len(arrayObjectSplit) > 1 {
				key = arrayObjectSplit[1]
			}

			if isArray {
				lineNum += len(array)
				object.Objects[key] = array
			} else {
				newObject := CreateYamlParser()
				err := newObject.createObject(buf[lineNum:], space+2)
				if err != nil {
					return errors.New("formatting error")
				}
				object.Objects[key] = newObject
			}
		}
	}
	return nil
}
