package parser

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func (ans Answer) Bool(def bool) (bool, error) {
	// if empty panic
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		return def, errors.New("no match")
	}

	// else convert (if incompatible strconv will panic)
	ret, err := strconv.ParseBool(temp)
	if err != nil {
		return def, errors.New("incompatible types")
	}
	return ret, nil
}

func (ans Answer) Int(def int) (int, error) {
	// if empty panic
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		return def, errors.New("no match")
	}

	// else convert (if incompatible strconv will panic)
	ret, err := strconv.Atoi(temp)
	if err != nil {
		return def, err
	}
	return ret, nil
}

func (ans Answer) String(def string) (ret string, err error) {
	ret = fmt.Sprintf("%v", ans.val)
	if ret == "<nil>" {
		return def, errors.New("no match")
	}

	// if not a string notify
	if reflect.TypeOf(ans.val) != reflect.TypeOf(ret) {
		fmt.Println("Mismatched types", reflect.TypeOf(ans.val), "and", reflect.TypeOf(ret))
		return def, errors.New("incompatible types")
	}
	return ret, nil
}

func (ans Answer) Float64(def float64) (float64, error) {
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		return def, errors.New("no match")
	}

	// else convert (if incompatible strconv will panic)
	ret, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		return def, err
	}
	return ret, nil
}

func (ans Answer) Duration(def time.Duration) (time.Duration, error) {
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		return def, errors.New("no match")
	}

	// else convert
	ret, err := time.ParseDuration(temp)
	if err != nil {
		return def, err
	}
	return ret, nil
}

func (ans Answer) StringSlice(def []string) ([]string, error) {
	var ret []string
	someInterface := new([]interface{})

	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		return def, errors.New("no match")
	}

	// if not a slice return
	if reflect.TypeOf(ans.val) != reflect.TypeOf(*someInterface) {
		return def, errors.New("mismatched types")
	}

	for i, val := range ans.val.([]interface{}) {
		// if not a string return
		if reflect.TypeOf(ans.val.([]interface{})[i]) != reflect.TypeOf(temp) {
			return def, errors.New("mismatched types")
		}
		// if string append
		ret = append(ret, val.(string))
	}

	return ret, nil
}

func (ans Answer) StringMap(def map[string]interface{}) (map[string]interface{}, error) {
	ret := new(map[string]interface{})

	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		return def, errors.New("no match")
	}

	// if not an Object return
	var tempObject Object

	if reflect.TypeOf(ans.val) != reflect.TypeOf(tempObject) {
		return def, errors.New("mismatched types")
	}

	// if Object save Object.Objects value
	tempObject = ans.val.(Object)
	*ret = tempObject.Objects

	return *ret, nil
}

func (ans Answer) Bytes() ([]byte, error) {
	// if empty panic
	temp := fmt.Sprintf("%v", ans.val)
	if temp == "<nil>" {
		return nil, errors.New("no match")
	}

	return []byte(temp), nil
}