package parser

import (
	"fmt"
	"reflect"
	"strconv"
)

type Mod interface {
	Bool(def bool) bool
	Int(def int) int
	String(def string) string
	/*	Float64(def float64) float64
		Duration(def time.Duration) time.Duration
		StringSlice(def []string) []string
		StringMap(def map[string]interface{}) map[string]interface{}
		Bytes() []byte
	*/
}

func (prop Property) Bool(def bool) bool {
	if prop.Mod != VAL_MOD {
		return def
	}

	value := fmt.Sprintf("%v", prop.Val)
	ans, err := strconv.ParseBool(value)
	if err != nil {
		return def
	}
	return ans
}

func (prop Property) Int(def int) int {
	if prop.Mod != VAL_MOD {
		return def
	}

	value := fmt.Sprintf("%v", prop.Val)
	ans, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return ans
}

func (prop Property) String(def string) string {
	ret := fmt.Sprintf("%v", prop.Val)
	if ret == "<nil>" {
		return def
	}
	// if not a string notify
	if reflect.TypeOf(prop.Val) != reflect.TypeOf(ret) {
		return def
	}
	return ret
}
