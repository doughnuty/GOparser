package parser

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

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

func (ans Answer) Float64() float64 {
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		panic("No match")
	}

	// else convert (if incompatible strconv will panic)
	ret, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return ret
}

func (ans Answer) Duration() time.Duration {
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		panic("No match")
	}

	// else convert
	ret, err := time.ParseDuration(temp)
	if err != nil {
		panic(err)
	}
	return ret
}

func (ans Answer) StringSlice() []string {
	var ret []string
	someInterface := new([]interface{})

	// if empty panic
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		panic("No match")
	}

	// if not a slice return
	if reflect.TypeOf(ans.val) != reflect.TypeOf(*someInterface) {
		fmt.Println("Mismatched types", reflect.TypeOf(ans.val), "and", reflect.TypeOf(*someInterface))
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

func (ans Answer) Bytes() []byte {
	// if empty panic
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		panic("No match")
	}

	return []byte(temp)
}