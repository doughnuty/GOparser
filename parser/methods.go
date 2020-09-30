package parser

import (
	"fmt"
	"strconv"
)

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
