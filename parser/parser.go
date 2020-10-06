package parser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"theRealParser/lexer"
	"theRealParser/token"
)

// true - is ok
// false - not ok
func (yaml *Yaml) checkIndentSpaces(spaces string) bool {
	valLen := len(spaces)
	return valLen > yaml.Spacing &&
		valLen%2 == 0
}

func NewYaml() Yaml {
	return Yaml{Map: make(map[string]Property), Spacing: 0}
}

func (yaml *Yaml) Parse(filename string) error {

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	str := string(buf)

	l := lexer.LexStart(filename, str)
	l.Adjacent = l.NextToken()

	err = yaml.parseTokens(l)
	if err != nil {
		return err
	}
	return err
}

func recursiveParse(l *lexer.Lexer) Yaml {
	// create new yaml
	newYaml := NewYaml()
	newYaml.Spacing = len(l.Current.Value)
	// write number of spaces and parse
	// newYaml.Spacing = countSpaces(l.adjacent.value)
	err := newYaml.parseTokens(l)
	if err != nil {
		return Yaml{
			Map:     nil,
			Spacing: 0,
		}
	}

	// assign new yaml to the Property
	return newYaml
}

func (yaml *Yaml) parseTokens(l *lexer.Lexer) error {

	keyVal := ""
	elemSlice := 0
	tempSlice := make([]string, 3, 5)

	for {
		l.Current = l.Adjacent
		l.Adjacent = l.NextToken()

		switch l.Current.Mod {
		case token.TOKEN_ERROR:
			return errors.New(l.Current.Value)

		case token.TOKEN_KEY:
			keyVal = l.Current.Value

		case token.TOKEN_VALUE:
			(*yaml).Map[keyVal] = Property{
				Mod: "value",
				Val: strings.TrimSpace(l.Current.Value),
			}
			keyVal = ""
			if len(l.Adjacent.Value) < yaml.Spacing {
				return nil
			}

		case token.TOKEN_DASH:
			tempSlice[elemSlice] = strings.TrimSpace(l.Current.Value)
			elemSlice++
			if len(l.Adjacent.Value) <= yaml.Spacing || l.Adjacent.Mod == token.TOKEN_EOF {
				(*yaml).Map[keyVal] = Property{
					Mod: ARR_MOD,
					Val: tempSlice,
				}
				return nil
			}
		case token.TOKEN_COLON:
			if l.Adjacent.Mod == token.TOKEN_SPACES && !yaml.checkIndentSpaces(l.Adjacent.Value) {
				// if colon followed by spaces and indentation is not proper report
				fmt.Println(l.Adjacent.Value)
				return errors.New("expected value found new line")
			}

		case token.TOKEN_SPACES:
			spaceNum := len(l.Current.Value)
			// if less spaces return
			if spaceNum < yaml.Spacing {
				//yaml.Spacing = countSpaces(token.value)
				return nil
			}
			// if more spaces and key create map
			if l.Adjacent.Mod == token.TOKEN_KEY && yaml.checkIndentSpaces(l.Current.Value) {
				(*yaml).Map[keyVal] = Property{
					Mod: MAP_MOD,
					Val: recursiveParse(l),
				}
				if l.Current.Mod == token.TOKEN_SPACES && len(l.Current.Value) < yaml.Spacing {
					return nil
				}
			}

		case token.TOKEN_EOF:
			return nil
		}
	}
}
