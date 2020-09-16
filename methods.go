package parser

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func (ans Answer) Bool(def bool) bool {
	// if empty panic
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		return def
	}

	// else convert (if incompatible strconv will panic)
	ret, err := strconv.ParseBool(temp)
	if err != nil {
		return def
	}
	return ret
}

func (ans Answer) Int(def int) int {
	// if empty panic
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		return def
	}

	// else convert (if incompatible strconv will panic)
	ret, err := strconv.Atoi(temp)
	if err != nil {
		return def
	}
	return ret
}

func (ans Answer) String(def string) (ret string) {
	ret = fmt.Sprintf("%v", ans.val)
	if ret == "<nil>" {
		return def
	}

	// if not a string notify
	if reflect.TypeOf(ans.val) != reflect.TypeOf(ret) {
		fmt.Println("Mismatched types", reflect.TypeOf(ans.val), "and", reflect.TypeOf(ret))
		return def
	}
	return ret
}

func (ans Answer) Float64(def float64) float64 {
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		return def
	}

	// else convert (if incompatible strconv will panic)
	ret, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		return def
	}
	return ret
}

func (ans Answer) Duration(def time.Duration) time.Duration {
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		return def
	}

	// else convert
	ret, err := time.ParseDuration(temp)
	if err != nil {
		return def
	}
	return ret
}

func (ans Answer) StringSlice(def []string) []string {
	var ret []string
	someInterface := new([]interface{})

	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		return def
	}

	// if not a slice return
	if reflect.TypeOf(ans.val) != reflect.TypeOf(*someInterface) {
		return def
	}

	for i, val := range ans.val.([]interface{}) {
		// if not a string return
		if reflect.TypeOf(ans.val.([]interface{})[i]) != reflect.TypeOf(temp) {
			return def
		}
		// if string append
		ret = append(ret, val.(string))
	}

	return ret
}

func (ans Answer) StringMap(def map[string]interface{}) map[string]interface{} {
	ret := new(map[string]interface{})

	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		return def
	}

	// if not an Object return
	var tempObject Object

	if reflect.TypeOf(ans.val) != reflect.TypeOf(tempObject) {
		return def
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
		return nil
	}

	return []byte(temp)
}
