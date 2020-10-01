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
func (yaml *Yaml) checkIndentSpaces(l *lexer.Lexer) bool {
	valLen := len(l.Adjacent.Value)
	return valLen > yaml.Spacing &&
		valLen%2 == 0
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

	l := lexer.LexStart(filename, str)
	l.Adjacent = l.NextToken()

	err = yaml.parseTokens(l)
	if err != nil {
		return err
	}
	return err
}

func (yaml *Yaml) parseTokens(l *lexer.Lexer) error {

	keyVal := ""

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

		case token.TOKEN_COLON:
			if l.Adjacent.Mod == token.TOKEN_SPACES && yaml.checkIndentSpaces(l) {
				// create new Property
				newProperty := Property{
					Mod: "map",
				}
				// create new yaml
				newYaml := NewYaml()

				// write number of spaces and parse
				// newYaml.Spacing = countSpaces(l.adjacent.value)
				err := newYaml.parseTokens(l)
				if err != nil {
					return err
				}

				// assign new yaml to the Property
				newProperty.Val = newYaml

				// add it to the last key value
				(*yaml).Map[keyVal] = newProperty
				if l.Current.Mod == token.TOKEN_SPACES && len(l.Current.Value) < yaml.Spacing {
					return nil
				}

			} else if l.Adjacent.Mod == token.TOKEN_SPACES && !yaml.checkIndentSpaces(l) {
				// if colon followed by spaces and indentation is not proper report
				fmt.Println(l.Adjacent.Value)
				return errors.New("expected value found new line")
			}
			/*else if l.Adjacent.Mod == token.TOKEN_VALUE {
				continue
			} else  {
				return errors.New("unexpected error")
			}*/

		case token.TOKEN_SPACES:
			spaceNum := len(l.Current.Value)
			if spaceNum < yaml.Spacing {
				//yaml.Spacing = countSpaces(token.value)
				return nil
			}
			yaml.Spacing = spaceNum
		case token.TOKEN_EOF:
			return nil
		}
	}
}
