package token

const COLON rune = ':'
const DASH rune = '-'
const EOF rune = 0
const HASH rune = '#'
const NL rune = '\n'
const SPACE rune = ' '

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
	TOKEN_DASH   // if dash sign
	TOKEN_NL     // if new line

	TOKEN_KEY
	TOKEN_VALUE
)
