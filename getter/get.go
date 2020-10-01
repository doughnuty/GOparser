package getter

import (
	"theRealParser/parser"
)

func (yaml *parser.Yaml) Get(path ...string) Mod {
	if yaml == nil {
		return parser.Property{Mod: "error"}
	}
	y := new(parser.Yaml)
	p := new(parser.Property)
	return p
}
