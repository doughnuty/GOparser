package token

const EOF 	rune = 0
const SPACE string = " "
const COLON string = ":"
const DASH 	string = "-"
const NL 	string = "\n"

type Token struct {
	Mod   TokenMod
	Value string
}

type TokenMod int

const (
	TOKEN_ERROR TokenMod = iota // if error occurred
	TOKEN_EOF                   // if end of the file

	TOKEN_SPACES					// if space
	TOKEN_COLON 					// if colon
	TOKEN_DASH 						// if dash sign
	TOKEN_NL 						// if new line

	TOKEN_KEY
	TOKEN_VALUE
	)

const (
	VAL_MOD string = "value"
	MAP_MOD string = "map"
	ARR_MOD string = "array"
)