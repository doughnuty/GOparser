package GOparser

import (
	"errors"
	"github.com/doughnuty/GOparser/lexer"
	"github.com/doughnuty/GOparser/token"
	"io/ioutil"
	"strings"
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
		//fmt.Println(err)
		return err
	}

	str := string(buf)

	l := lexer.LexStart(str)
	l.Following = l.NextToken()

	err = yaml.parseTokens(l)
	if err != nil {
		return err
	}
	return err
}

func (yaml *Yaml) recursiveParse(l *lexer.Lexer) error {
	// create new yaml
	*yaml = NewYaml()
	yaml.Spacing = len(l.Current.Value)
	// write number of spaces and parse
	// yaml.Spacing = countSpaces(l.adjacent.value)
	err := yaml.parseTokens(l)
	if err != nil {
		return err
	}

	// assign new yaml to the Property
	return nil
}

func (yaml *Yaml) parseTokens(l *lexer.Lexer) error {

	keyVal := ""
	tempSlice := make([]string, 0, 10)

	for {
		l.Current = l.Following
		l.Following = l.NextToken()

		switch l.Current.Mod {
		case token.TOKEN_ERROR:
			return errors.New(l.Current.Value)

		case token.TOKEN_KEY:
			keyVal = strings.TrimSpace(l.Current.Value)

		case token.TOKEN_VALUE:
			(*yaml).Map[keyVal] = Property{
				Mod: "value",
				Val: strings.TrimSpace(l.Current.Value),
			}
			keyVal = ""
			if len(l.Following.Value) < yaml.Spacing || l.Following.Mod == token.TOKEN_EOF {
				return nil
			}

		case token.TOKEN_ARRAY:
			tempSlice = append(tempSlice, strings.TrimSpace(l.Current.Value))
			if len(l.Following.Value) <= yaml.Spacing || l.Following.Mod == token.TOKEN_EOF {
				(*yaml).Map[keyVal] = Property{
					Mod: ARR_MOD,
					Val: tempSlice,
				}
			}
		case token.TOKEN_COLON:
			if l.Following.Mod == token.TOKEN_SPACES && !yaml.checkIndentSpaces(l.Following.Value) {
				// if colon followed by spaces and indentation is not proper report
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
			if l.Following.Mod == token.TOKEN_KEY {
				if yaml.checkIndentSpaces(l.Current.Value) && keyVal != "" {
					var newYaml Yaml
					err := newYaml.recursiveParse(l)
					if err != nil {
						return err
					}
					(*yaml).Map[keyVal] = Property{
						Mod: MAP_MOD,
						Val: newYaml,
					}
					//yaml.Spacing = spaceNum
					if l.Current.Mod == token.TOKEN_SPACES && len(l.Current.Value) < yaml.Spacing {
						return nil
					} else if l.Following.Mod == token.TOKEN_SPACES && len(l.Following.Value) < yaml.Spacing {
						return nil
					}
				}

			}

		case token.TOKEN_EOF:
			return nil
		}
		//
	}
}
