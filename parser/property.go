package parser

import (
	"errors"
	"strings"
	"theRealParser/lexer"
	"theRealParser/token"
)

func (yaml *Yaml) parseTokens(lexer *lexer.Lexer) error {

	keyVal := ""

	for {
		lexer.Current = lexer.Adjacent
		lexer.Adjacent = lexer.NextToken()

		switch lexer.Current.Mod {
		case token.TOKEN_ERROR:
			return errors.New(lexer.Current.Value)

		case token.TOKEN_KEY:
			keyVal = lexer.Current.Value

		case token.TOKEN_VALUE:
			(*yaml).Map[keyVal] = Property{
				mod: "value",
				val: strings.TrimSpace(lexer.Current.Value),
			}
			keyVal = ""

		case token.TOKEN_COLON:
			if lexer.Adjacent.Mod == token.TOKEN_SPACES {
				// create new Property
				newProperty := Property{
					mod: "map",
				}
				// create new yaml
				newYaml := NewYaml()

				// write number of spaces and parse
				// newYaml.Spacing = countSpaces(lexer.adjacent.value)
				err := newYaml.parseTokens(lexer)
				if err != nil {
					return err
				}

				// assign new yaml to the Property
				newProperty.val = newYaml

				// add it to the last key value
				(*yaml).Map[keyVal] = newProperty
				if lexer.Current.Mod == token.TOKEN_SPACES && len(lexer.Current.Value) < yaml.Spacing {
					return nil
				}

			} else if lexer.Adjacent.Mod == token.TOKEN_VALUE {
				continue
			} else {
				return errors.New("unexpected colon")
			}

		case token.TOKEN_SPACES:
			spaceNum := len(lexer.Current.Value)
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
