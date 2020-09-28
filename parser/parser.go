package parser

import (
	"fmt"
	"io/ioutil"
	lexer2 "theRealParser/lexer"
	"time"
)

type Yaml struct {
	Map map[string]property
	Spacing int
}

type property struct {
	mod string
	val interface{}
}

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
	return Yaml{Map: make(map[string]property)}
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
	fmt.Println(lexer.Input)
	err = yaml.parseTokens(lexer)
	if err != nil {
		return err
	}
	return err
}

func (yaml *Yaml) PrintYaml() {
	fmt.Println(yaml)
}

/*
func (yaml Yaml) Get(path ...string) Mod {
	if yaml == nil {

	}
}
*/
