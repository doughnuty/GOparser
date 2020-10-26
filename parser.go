package GOparser

import (
	"errors"
	lexer "github.com/doughnuty/GOparser/lexer_no_chan"
	"github.com/doughnuty/GOparser/token"
	"io/ioutil"
	"strings"
)

func pop(yamlSlice []Yaml) (Yaml, []Yaml) {
	return yamlSlice[len(yamlSlice)-1], yamlSlice[:len(yamlSlice)-1]
}

// true - is ok
// false - not ok
func (yaml *Yaml) checkIndentSpaces(spaces string) bool {
	valLen := len(spaces)
	return valLen > yaml.Spacing &&
		valLen%2 == 0
}

func (yaml *Yaml) ParseFiles(filename string) error {

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

func newString(value string) (string, error) {
	runVal := []rune(value)
	if len(runVal) == 0 {
		return value, nil
	}
	newVal := make([]rune, len(value))
	j := 0
	isEscape := false
	if runVal[0] == '"' {
		for i, c := range value {
			if isEscape {
				switch c {
				case 't':
					c = '\t'
				case 'b':
					c = '\b'
				case 'n':
					c = '\n'
				case 'r':
					c = '\r'
				case 'f':
					c = '\f'
				case '"':
					c = '"'
				case '\\':
					c = '\\'
				default:
					return value, errors.New("unknown escape")
				}
				isEscape = false
			} else if c == '\\' {
				isEscape = true
				continue
			} else if c == '"' {
				if i == len(value)-1 {
					newVal = newVal[:j]
					return string(newVal), nil
				} else if i == 0 {
					continue
				} else {
					return value, errors.New("unexpected characters after quote")
				}
			}

			newVal[j] = c
			j++
		}
		return value, errors.New("missing quote at the end of string")
	}
	return value, nil
}

func (yaml *Yaml) parseTokens(l *lexer.Lexer) error {

	keyVal := ""
	tempSlice := make([]string, 0, 10)
	yamlSlice := make([]Yaml, 0, 10)
	for {
		l.Current = l.Following
		l.Following = l.NextToken()

		switch l.Current.Mod {
		case token.TOKEN_ERROR:
			return errors.New(l.Current.Value)

		case token.TOKEN_KEY:
			keyVal = strings.TrimSpace(l.Current.Value)

		case token.TOKEN_VALUE:
			value, err := newString(strings.TrimSpace(l.Current.Value))
			if err != nil {
				return err
			}
			(*yaml).Map[keyVal] = Property{
				Mod: VAL_MOD,
				Val: value,
			}
			keyVal = ""
			if len(l.Following.Value) < yaml.Spacing || l.Following.Mod == token.TOKEN_EOF {
				// pop yaml
				*yaml, yamlSlice = pop(yamlSlice)
			}

		case token.TOKEN_ARRAY:
			value, err := newString(strings.TrimSpace(l.Current.Value))
			if err != nil {
				return err
			}
			tempSlice = append(tempSlice, value)
			if len(l.Following.Value) <= yaml.Spacing || l.Following.Mod == token.TOKEN_EOF {
				(*yaml).Map[keyVal] = Property{
					Mod: ARR_MOD,
					Val: tempSlice,
				}
				tempSlice = make([]string, 0, 10)
				*yaml, yamlSlice = pop(yamlSlice)
			}
		case token.TOKEN_COLON:
			if l.Following.Mod == token.TOKEN_SPACES && !yaml.checkIndentSpaces(l.Following.Value) {
				// if colon followed by spaces and indentation is not proper report
				return errors.New("expected value found new line")
			}

		case token.TOKEN_SPACES:
			spaceNum := len(l.Current.Value)

			// if less spaces pop
			for spaceNum < yaml.Spacing {
				*yaml, yamlSlice = pop(yamlSlice)
			}
			/*if spaceNum < yaml.Spacing {
				*yaml, yamlSlice = pop(yamlSlice)
			}*/
			// if more spaces and key create map
			if l.Following.Mod == token.TOKEN_KEY {
				if yaml.checkIndentSpaces(l.Current.Value) && keyVal != "" {
					oldYaml := *yaml
					yamlSlice = append(yamlSlice, oldYaml)
					newYaml := NewYaml()
					newYaml.Spacing = len(l.Current.Value)
					*yaml = newYaml
					oldYaml.Map[keyVal] = Property{
						Mod: MAP_MOD,
						Val: *yaml,
					}
					keyVal = ""
				}
			}

		case token.TOKEN_EOF:
			return nil
		}
	}
}
