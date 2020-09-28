package lexer

import (
	"strings"
	"theRealParser/errors"
	"theRealParser/token"
	"unicode"
)

type lexState func(*Lexer) lexState

// place error into chan
func (lexer *Lexer) error(format string) lexState {
	lexer.tokens <- token.Token{
		Mod:   token.TOKEN_ERROR,
		Value: format,
	}
	return nil
}

func lexValue(lexer *Lexer) lexState {
	for {
		if unicode.IsSpace(lexer.next()) {
			lexer.pos--
			lexer.putToken(token.TOKEN_VALUE)
			lexer.skipBlank()
			return lexBegin
		}

		if lexer.isEOF() {
			//return lexer.error(errors.LEXER_ERROR_UNEXPECTED_EOF)
			lexer.putToken(token.TOKEN_VALUE)
			return lexEOF
		}

		lexer.increment()

	}
}

func lexColumn(lexer *Lexer) lexState {
	lexer.pos += len(token.COLON)
	lexer.putToken(token.TOKEN_COLON)

	lexer.skipBlank()
	if strings.HasPrefix(lexer.toEnd(), token.NL) {
		lexer.increment()
		lexer.ignore()
		return lexBegin
	}
	return lexValue
}

func lexEOF(lexer *Lexer) lexState {
	lexer.putToken(token.TOKEN_EOF)
	return nil
}
func lexKey(lexer *Lexer) lexState {
	for {
		if lexer.isEOF() {
			return lexer.error(errors.LEXER_ERROR_MISSING_COLON)
		}

		if strings.HasPrefix(lexer.toEnd(), token.COLON) {
			lexer.putToken(token.TOKEN_KEY)
			return lexColumn
		}

		lexer.increment()
	}
}

func lexIndent(lexer *Lexer) lexState {
	t := rune(lexer.Input[lexer.pos])
	for unicode.IsSpace(t) {
		if t != ' ' {
			lexer.ignore()
		}
		t = lexer.next()
	}
	lexer.pos--
	if t == token.EOF {
		return lexEOF
	} else {
		lexer.putToken(token.TOKEN_SPACES)
	}
	return lexKey
}

func lexBegin(lexer *Lexer) lexState {
	t := rune(lexer.Input[lexer.pos])
	if unicode.IsSpace(t) {
		lexer.pos--
		return lexIndent
	}
	return lexKey
}

func LexerStart(name, input string) *Lexer {
	lexer := &Lexer{
		name:   name,
		Input:  input,
		state:  lexBegin,
		tokens: make(chan token.Token, 3),
	}

	return lexer
}
