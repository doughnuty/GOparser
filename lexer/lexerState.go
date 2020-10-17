package lexer

import (
	"fmt"
	"github.com/doughnuty/GOparser/errors"
	"github.com/doughnuty/GOparser/token"
	"strings"
	"unicode"
)

type lexState func(*Lexer) lexState

func LexStart(input string) *Lexer {
	lexer := &Lexer{
		input:  input,
		state:  lexBegin,
		tokens: make(chan token.Token, 3),
	}

	return lexer
}

func lexBegin(lexer *Lexer) lexState {
	t := rune(lexer.input[lexer.pos])
	if unicode.IsSpace(t) {
		lexer.pos--
		return lexIndent
	}
	if t == token.HASH {
		return lexComment
	}
	return lexKey
}

// place error into chan
func (lexer *Lexer) error(format string) lexState {
	lexer.tokens <- token.Token{
		Mod:   token.TOKEN_ERROR,
		Value: format,
	}
	return lexEOF
}

func lexComment(lexer *Lexer) lexState {
	for {
		if lexer.isEOF() {
			lexer.ignore()
			return lexEOF
		}

		if strings.HasPrefix(lexer.toEnd(), string(token.NL)) {
			lexer.ignore()
			return lexIndent
		}

		lexer.increment()
	}
}

// Возможно стоит сделать так что бегин возвращает ТОЛЬКО индент
// чтоб не дублировать один и тот же код для этих функций
// Либо в целом отказаться от функции бегин чтоб не засорять
// Начинать сразу с индента

// Additionally, I can change it completely by creating one "factory" function
// it will loop through input values deciding which lex function to run
// This way, code will be easier to read (I guess)
// Not so good idea tho
func lexIndent(lexer *Lexer) lexState {
	if lexer.isEOF() {
		return lexEOF
	}

	t := lexer.next()
	for unicode.IsSpace(t) {
		if t != token.SPACE {
			lexer.ignore()
		}
		t = lexer.next()
	}
	//fmt.Println(string(lexer.input[lexer.pos]), lexer.pos)
	lexer.pos--
	//fmt.Println(string(lexer.input[lexer.pos]), lexer.pos)
	if t == token.EOF {
		return lexEOF
	}

	if t == token.HASH {
		return lexComment
	}

	if (lexer.pos-lexer.start)%2 != 0 {
		return lexer.error(errors.LEXER_BAD_INDENTATION)
	}
	lexer.putToken(token.TOKEN_SPACES)
	//	lexer.increment()

	if t == token.DASH {
		return lexArrayDash
	}

	return lexKey
}

func lexArrayDash(lexer *Lexer) lexState {
	for {
		if lexer.isEOF() {
			lexer.pos--
			lexer.putToken(token.TOKEN_ARRAY)
			return lexEOF
		}
		if strings.HasPrefix(lexer.toEnd(), string(token.DASH)) {
			lexer.increment()
			lexer.ignore()
		}

		if strings.HasPrefix(lexer.toEnd(), string(token.HASH)) {
			lexer.putToken(token.TOKEN_ARRAY)
			return lexComment
		}

		if strings.HasPrefix(lexer.toEnd(), string(token.NL)) {
			lexer.putToken(token.TOKEN_ARRAY)
			return lexIndent
		}

		lexer.increment()
	}
}

func lexArrayBracket(lexer *Lexer) lexState {
	for {
		if lexer.isEOF() {
			lexer.pos--
			lexer.putToken(token.TOKEN_ARRAY)
			return lexEOF
		}

		switch rune(lexer.input[lexer.pos]) {
		case token.LBRACKET:
			lexer.increment()
			lexer.ignore()
		case token.COMMA:
			lexer.putToken(token.TOKEN_ARRAY)
			lexer.increment()
			lexer.ignore()
			return lexArrayBracket
		case token.RBRACKET:
			lexer.putToken(token.TOKEN_ARRAY)
			lexer.pos++
			if lexer.skipBlank() {
				return lexEOF
			} else if strings.HasPrefix(lexer.toEnd(), string(token.NL)) {
				return lexIndent
			} else if strings.HasPrefix(lexer.toEnd(), string(token.HASH)) {
				return lexComment
			} else {
				return lexer.error(errors.LEXER_BAD_INDENTATION)
			}
		case token.NL:
			return lexer.error(errors.LEXER_MISSING_BRACKET)
		}

		lexer.increment()
	}
}

func lexKey(lexer *Lexer) lexState {
	for {
		if lexer.isEOF() {
			return lexer.error(errors.LEXER_ERROR_MISSING_COLON)
		}

		if strings.HasPrefix(lexer.toEnd(), string(token.NL)) {
			fmt.Println(lexer.input[lexer.pos-2])
			return lexer.error(errors.LEXER_ERROR_MISSING_COLON)
		}

		if strings.HasPrefix(lexer.toEnd(), string(token.COLON)) {
			lexer.putToken(token.TOKEN_KEY)
			return lexColumn
		}

		lexer.increment()
	}
}

func lexColumn(lexer *Lexer) lexState {
	lexer.pos += 1
	lexer.putToken(token.TOKEN_COLON)

	// skip blank function return whether the file ended
	// true for EOF false for not EOF
	if lexer.skipBlank() {
		return lexer.error(errors.LEXER_ERROR_UNEXPECTED_EOF)
	}
	if strings.HasPrefix(lexer.toEnd(), string(token.LBRACKET)) {
		return lexArrayBracket
	}

	if strings.HasPrefix(lexer.toEnd(), string(token.NL)) {
		lexer.increment()
		if lexer.isEOF() {
			lexer.pos--
			return lexEOF
		}
		lexer.ignore()
		return lexIndent
	}
	return lexValue
}

func lexValue(lexer *Lexer) lexState {

	if rune(lexer.input[lexer.pos]) == token.DASH {
		return lexer.error(errors.LEXER_BAD_INDENTATION)
	}

	for {
		if lexer.isEOF() {
			lexer.pos--
			lexer.putToken(token.TOKEN_VALUE)
			return lexEOF
		}

		if strings.HasPrefix(lexer.toEnd(), string(token.HASH)) {
			lexer.putToken(token.TOKEN_VALUE)

			return lexComment
		}

		if strings.HasPrefix(lexer.toEnd(), string(token.NL)) {
			lexer.putToken(token.TOKEN_VALUE)
			if lexer.skipBlank() {
				return lexEOF
			}

			return lexIndent
		}

		if strings.HasPrefix(lexer.toEnd(), string(token.COLON)) {
			return lexer.error(errors.LEXER_BAD_INDENTATION)
		}

		lexer.increment()
	}
}

func lexEOF(lexer *Lexer) lexState {
	if lexer.pos < lexer.start {
		lexer.start = lexer.pos
	}
	lexer.putToken(token.TOKEN_EOF)
	return nil
}
