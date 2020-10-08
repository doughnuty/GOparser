package parser

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type Mod interface {
	Bool(def bool) bool
	Int(def int) int
	String(def string) string
	Float64(def float64) float64
	Duration(def time.Duration) time.Duration
	StringSlice(def []string) []string
	StringMap(def map[string]string) map[string]string
	Bytes() []byte
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
	if prop.Mod != VAL_MOD {
		return def
	}
	ret := fmt.Sprintf("%v", prop.Val)
	// if not a string notify
	if reflect.TypeOf(prop.Val) != reflect.TypeOf(ret) {
		return def
	}
	return ret
}

func (prop Property) Float64(def float64) float64 {
	if prop.Mod != VAL_MOD {
		return def
	}

	temp := fmt.Sprintf("%v", prop.Val)

	// else convert
	ret, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		return def
	}
	return ret
}

func (prop Property) Duration(def time.Duration) time.Duration {
	if prop.Mod != VAL_MOD {
		return def
	}
	temp := fmt.Sprintf("%v", prop.Val)
	ret, err := time.ParseDuration(temp)
	if err != nil {
		return def
	}
	return ret
}

func (prop Property) StringSlice(def []string) []string {
	if prop.Mod != ARR_MOD {
		return def
	}

	ret := prop.Val.([]string)
	return ret
}

func (prop Property) StringMap(def map[string]string) map[string]string {
	if prop.Mod != MAP_MOD {
		return def
	}
	ret := make(map[string]string, 10)
	for i, j := range prop.Val.(Yaml).Map {
		if j.Mod != VAL_MOD {
			return def
		}
		ret[i] = j.Val.(string)
	}
	return ret
}

func (prop Property) Bytes() []byte {
	b := []byte(fmt.Sprintf("%v", prop.Val))
	return b
}
