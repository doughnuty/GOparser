package parser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"theRealParser/lexer"
	"theRealParser/token"
)

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

	l := lexer.LexerStart(filename, str)
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
			if l.Adjacent.Mod == token.TOKEN_SPACES {
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

			} else if l.Adjacent.Mod == token.TOKEN_VALUE {
				continue
			} else {
				return errors.New("unexpected colon")
			}

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
