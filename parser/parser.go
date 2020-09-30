package parser

import (
	"fmt"
	"io/ioutil"
	lexer2 "theRealParser/lexer"
	"time"
)

type Yaml struct {
	Map     map[string]Property
	Spacing int
}

type Property struct {
	mod string
	val interface{}
}

const (
	VAL_MOD string = "value"
	MAP_MOD string = "map"
	ARR_MOD string = "array"
)

type Mod interface {
	Bool(def bool) bool
	Int(def int) int
	String(def string) string
	Float64(def float64) float64
	Duration(def time.Duration) time.Duration
	StringSlice(def []string) []string
	StringMap(def map[string]interface{}) map[string]interface{}
	Bytes() []byte
}

func NewYaml() Yaml {
	return Yaml{Map: make(map[string]Property)}
}

func (yaml *Yaml) Parse(filename string) error {

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	str := string(buf)

	lexer := lexer2.LexerStart(filename, str)
	lexer.Adjacent = lexer.NextToken()

	err = yaml.parseTokens(lexer)
	if err != nil {
		return err
	}
	return err
}

/*
func (yaml Yaml) Get(path ...string) Mod {
	if yaml == nil {
		return nil
	}

	return Property
}
*/
