package parser

import (
	"reflect"
)

// Searches for values from path in the source Object
func (ans *Answer) seekVal(source Object, path []string) interface{} {
	//	fmt.Println(source.Objects[path[0]])
	// how much elements in path
	length := len(path)
	// loop through
	for pathNum, i := range path {
		// if its last element in path return
		if pathNum == length - 1 {
			return source.Objects[path[length-1]]
		}
		// if type of value is Object update source
		temp := source.Objects[i]
		var someObject Object
		someArray := new([]interface{})
		if reflect.TypeOf(*someArray) == reflect.TypeOf(temp) {
			*someArray = temp.([]interface{})
			pathNum++
			for j := range *someArray {
				someObject = (*someArray)[j].(Object)
				if someObject.Objects[path[pathNum]] != nil {
					source = someObject
				}
			}
		} else if reflect.TypeOf(temp) == reflect.TypeOf(someObject) {
			someObject = temp.(Object)
			source = someObject
		}
	}
	return source.Objects[path[length-1]]
}

