package lexer

import (
	"github.com/doughnuty/GOparser/token"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input  string
	tokens chan token.Token
	state  lexState

	start int
	pos   int
	width int

	Following token.Token
	Current   token.Token
}

func (lexer *Lexer) ignore() {
	lexer.start = lexer.pos
}

// increment position
func (lexer *Lexer) increment() {
	lexer.pos++
}

// check if EOF
func (lexer *Lexer) isEOF() bool {
	//fmt.Println(lexer.input)
	return lexer.pos >= len(lexer.input)
}

// advances to the next position
func (lexer *Lexer) next() rune {
	if lexer.pos >= utf8.RuneCountInString(lexer.input) {
		lexer.width = 0
		return token.EOF
	}

	result, width := utf8.DecodeRuneInString(lexer.input[lexer.pos:])

	lexer.width = width
	lexer.pos += lexer.width
	return result
}

// return next token from channel
func (lexer *Lexer) NextToken() token.Token {
	// if we reached the end no reason to search for new token
	if lexer.Current.Mod == token.TOKEN_EOF {
		return lexer.Current
	}

	for {
		select {
		case t := <-lexer.tokens:
			return t
		default:
			lexer.state = lexer.state(lexer)
		}
	}
}

// put token into token channel
func (lexer *Lexer) putToken(tokenMod token.TokenMod) {
	lexer.tokens <- token.Token{Mod: tokenMod, Value: lexer.input[lexer.start:lexer.pos]}
	lexer.start = lexer.pos
}

// skip spaces until something meaningful or a new line appears
func (lexer *Lexer) skipBlank() (isEOF bool) {
	for {
		ch := lexer.next()

		if ch == token.EOF {
			//lexer.putToken(token.TOKEN_EOF)
			isEOF = true
			break
		}

		if ch == token.NL || !unicode.IsSpace(ch) {
			lexer.pos--
			lexer.ignore()
			isEOF = false
			break
		}

	}
	return
}

func (lexer *Lexer) skipLine() {
	for {
		ch := lexer.next()

		if ch == '\n' {
			lexer.ignore()
			break
		}

		if ch == token.EOF {
			lexer.putToken(token.TOKEN_EOF)
			break
		}
	}
}

func (lexer *Lexer) toEnd() string {
	return lexer.input[lexer.pos:]
}
