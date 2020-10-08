package token

const COLON rune = ':'
const DASH rune = '-'
const EOF rune = 0
const HASH rune = '#'
const NL rune = '\n'
const SPACE rune = ' '
const LBRACKET rune = '['
const RBRACKET rune = ']'
const COMMA rune = ','

type Token struct {
	Mod   TokenMod
	Value string
}

type TokenMod int

const (
	TOKEN_ERROR TokenMod = iota // if error occurred
	TOKEN_EOF                   // if end of the file

	TOKEN_SPACES // if space
	TOKEN_COLON  // if colon
	TOKEN_ARRAY  // if array

	TOKEN_KEY
	TOKEN_VALUE
)
