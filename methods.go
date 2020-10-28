package GOparser

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
	if prop.mod != VAL_MOD {
		return def
	}

	value := fmt.Sprintf("%v", prop.val)
	ans, err := strconv.ParseBool(value)
	if err != nil {
		return def
	}
	return ans
}

func (prop Property) Int(def int) int {
	if prop.mod != VAL_MOD {
		return def
	}

	value := fmt.Sprintf("%v", prop.val)
	ans, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return ans
}

func (prop Property) String(def string) string {
	if prop.mod != VAL_MOD {
		return def
	}
	ret := fmt.Sprintf("%v", prop.val)
	// if not a string notify
	if reflect.TypeOf(prop.val) != reflect.TypeOf(ret) {
		return def
	}
	return ret
}

func (prop Property) Float64(def float64) float64 {
	if prop.mod != VAL_MOD {
		return def
	}

	temp := fmt.Sprintf("%v", prop.val)

	// else convert
	ret, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		return def
	}
	return ret
}

func (prop Property) Duration(def time.Duration) time.Duration {
	if prop.mod != VAL_MOD {
		return def
	}
	temp := fmt.Sprintf("%v", prop.val)
	ret, err := time.ParseDuration(temp)
	if err != nil {
		return def
	}
	return ret
}

func (prop Property) StringSlice(def []string) []string {
	if prop.mod != ARR_MOD {
		return def
	}

	ret := prop.val.([]string)
	return ret
}

func (prop Property) StringMap(def map[string]string) map[string]string {
	if prop.mod != MAP_MOD {
		return def
	}
	ret := make(map[string]string, 10) // edit to handle more data
	for i, j := range prop.val.(Yaml).Map {
		if j.mod != VAL_MOD {
			return def
		}
		ret[i] = j.val.(string)
	}
	return ret
}

func (prop Property) Bytes() []byte {
	b := []byte(fmt.Sprintf("%v", prop.val))
	return b
}
