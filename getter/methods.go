package getter

import (
	"fmt"
	"strconv"
	"theRealParser/parser"
	"time"
)

type Mod interface {
	Bool(def bool) bool
	Int(def int) int
	/*	String(def string) string
		Float64(def float64) float64
		Duration(def time.Duration) time.Duration
		StringSlice(def []string) []string
		StringMap(def map[string]interface{}) map[string]interface{}
		Bytes() []byte
	*/
}

func (prop parser.Property) Bool(def bool) bool {
	if prop.Mod != parser.VAL_MOD {
		return def
	}

	value := fmt.Sprintf("%v", prop.val)
	ans, err := strconv.ParseBool(value)
	if err != nil {
		return def
	}
	return ans
}

func (prop parser.Property) Int(def int) int {
	if prop.Mod != parser.VAL_MOD {
		return def
	}

	value := fmt.Sprintf("%v", prop.val)
	ans, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return ans
}
