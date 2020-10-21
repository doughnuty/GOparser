package lexer_no_chan

import (
	"github.com/doughnuty/GOparser/errors"
	"github.com/doughnuty/GOparser/token"
	"strings"
	"unicode"
)

type lexState func(*Lexer) token.Token

func LexStart(input string) *Lexer {
	lexer := &Lexer{
		input: input,
		state: lexIndent,
	}

	return lexer
}

// place error into chan
func (lexer *Lexer) error(format string) token.Token {
	result := token.Token{
		Mod:   token.TOKEN_ERROR,
		Value: format,
	}

	lexer.changeState(lexEOF)

	return result
}

/*func lexComment(lexer *Lexer) token.Token {
	lexer.skipLine()
}
*/
// Возможно стоит сделать так что бегин возвращает ТОЛЬКО индент
// чтоб не дублировать один и тот же код для этих функций
// Либо в целом отказаться от функции бегин чтоб не засорять
// Начинать сразу с индента

// Additionally, I can change it completely by creating one "factory" function
// it will loop through input values deciding which lex function to run
// This way, code will be easier to read (I guess)
// Not so good idea tho
func lexIndent(lexer *Lexer) token.Token {
	if lexer.isEOF() {
		return lexer.putToken(token.TOKEN_EOF)
	}

	t := lexer.next()
	for unicode.IsSpace(t) {
		if t != token.SPACE {
			lexer.ignore()
		}
		t = lexer.next()
	}
	lexer.pos--

	if t == token.EOF {
		return lexer.putToken(token.TOKEN_EOF)
	}

	if t == token.HASH {
		lexer.skipLine()
		lexer.changeState(lexIndent)
		return lexer.state(lexer)
	}

	if (lexer.pos-lexer.start)%2 != 0 {
		return lexer.error(errors.LEXER_BAD_INDENTATION)
	}

	if t == token.DASH {
		lexer.changeState(lexArrayDash)
	} else {
		lexer.changeState(lexKey)
	}

	return lexer.putToken(token.TOKEN_SPACES)
}

func lexArrayDash(lexer *Lexer) token.Token {
	for {
		if lexer.isEOF() {
			lexer.pos--
			lexer.state = lexEOF
			return lexer.putToken(token.TOKEN_ARRAY)
		}
		if strings.HasPrefix(lexer.toEnd(), string(token.DASH)) {
			lexer.increment()
			lexer.ignore()
		}

		if strings.HasPrefix(lexer.toEnd(), string(token.HASH)) {
			result := lexer.putToken(token.TOKEN_ARRAY)
			lexer.skipLine()
			lexer.changeState(lexIndent)
			return result
		}

		if strings.HasPrefix(lexer.toEnd(), string(token.NL)) {
			lexer.changeState(lexIndent)
			return lexer.putToken(token.TOKEN_ARRAY)
		}

		lexer.increment()
	}
}

func lexArrayBracket(lexer *Lexer) token.Token {
	for {
		if lexer.isEOF() {
			lexer.pos--
			return lexer.putToken(token.TOKEN_ARRAY)
		}

		switch rune(lexer.input[lexer.pos]) {
		case token.LBRACKET:
			lexer.increment()
			lexer.ignore()
		case token.COMMA:
			result := lexer.putToken(token.TOKEN_ARRAY)
			lexer.increment()
			lexer.ignore()
			lexer.state = lexArrayBracket
			return result
		case token.RBRACKET:
			result := lexer.putToken(token.TOKEN_ARRAY)
			lexer.increment()
			if lexer.skipBlank() {
				lexer.changeState(lexEOF)
			} else if strings.HasPrefix(lexer.toEnd(), string(token.NL)) {
				lexer.changeState(lexIndent)
			} else if strings.HasPrefix(lexer.toEnd(), string(token.HASH)) {
				lexer.skipLine()
				lexer.changeState(lexIndent)
			} else {
				return lexer.error(errors.LEXER_BAD_INDENTATION)
			}
			return result
		case token.NL:
			return lexer.error(errors.LEXER_MISSING_BRACKET)
		}

		lexer.increment()
	}
}

func lexKey(lexer *Lexer) token.Token {
	for {
		if lexer.isEOF() {
			return lexer.error(errors.LEXER_ERROR_MISSING_COLON)
		}

		if strings.HasPrefix(lexer.toEnd(), string(token.NL)) {
			return lexer.error(errors.LEXER_ERROR_MISSING_COLON)
		}

		if strings.HasPrefix(lexer.toEnd(), string(token.COLON)) {
			lexer.changeState(lexColumn)
			return lexer.putToken(token.TOKEN_KEY)
		}

		lexer.increment()
	}
}

func lexColumn(lexer *Lexer) token.Token {
	lexer.increment()
	result := lexer.putToken(token.TOKEN_COLON)

	// skip blank function return whether the file ended
	// true for EOF false for not EOF
	if lexer.skipBlank() {
		return lexer.error(errors.LEXER_ERROR_UNEXPECTED_EOF)
	}
	if strings.HasPrefix(lexer.toEnd(), string(token.LBRACKET)) {
		lexer.changeState(lexArrayBracket)
	} else if strings.HasPrefix(lexer.toEnd(), string(token.NL)) {
		lexer.increment()
		if lexer.isEOF() {
			lexer.pos--
			lexer.state = lexEOF
		} else {
			lexer.ignore()
			lexer.state = lexIndent
		}
	} else {
		lexer.state = lexValue
	}

	return result
}

func lexValue(lexer *Lexer) token.Token {
	var result token.Token
	if rune(lexer.input[lexer.pos]) == token.DASH {
		return lexer.error(errors.LEXER_BAD_INDENTATION)
	}

	for {
		if lexer.isEOF() {
			lexer.pos--
			lexer.changeState(lexEOF)
			return lexer.putToken(token.TOKEN_VALUE)
		}

		if strings.HasPrefix(lexer.toEnd(), string(token.HASH)) {
			result = lexer.putToken(token.TOKEN_VALUE)
			lexer.skipLine()
			lexer.changeState(lexIndent)
			return result
		}

		if strings.HasPrefix(lexer.toEnd(), string(token.NL)) {
			result = lexer.putToken(token.TOKEN_VALUE)
			if lexer.skipBlank() {
				lexer.changeState(lexEOF)
			} else {
				lexer.changeState(lexIndent)
			}
			return result
		}

		if strings.HasPrefix(lexer.toEnd(), string(token.COLON)) {
			return lexer.error(errors.LEXER_BAD_INDENTATION)
		}

		lexer.increment()
	}
}

func lexEOF(lexer *Lexer) token.Token {
	if lexer.pos < lexer.start {
		lexer.start = lexer.pos
	}
	lexer.state = nil
	return lexer.putToken(token.TOKEN_EOF)
}
