package lexer

import (
	"strings"
	"theRealParser/errors"
	"theRealParser/token"
	"unicode"
)

type lexState func(*Lexer) lexState

func LexStart(name, input string) *Lexer {
	lexer := &Lexer{
		name:   name,
		Input:  input,
		state:  lexBegin,
		tokens: make(chan token.Token, 3),
	}

	return lexer
}

func lexBegin(lexer *Lexer) lexState {
	t := rune(lexer.Input[lexer.pos])
	if unicode.IsSpace(t) {
		lexer.pos--
		return lexIndent
	} else if t == token.HASH {

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

func lexIndent(lexer *Lexer) lexState {
	t := rune(lexer.Input[lexer.pos])
	for unicode.IsSpace(t) {
		if t != token.SPACE {
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

func lexKey(lexer *Lexer) lexState {
	for {
		if lexer.isEOF() {
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

	if lexer.skipBlank() {
		return lexer.error(errors.LEXER_ERROR_UNEXPECTED_EOF)
	}
	if strings.HasPrefix(lexer.toEnd(), string(token.NL)) {
		lexer.increment()
		if lexer.isEOF() {
			lexer.pos--
			return lexEOF
		}
		lexer.ignore()
		return lexBegin
	}
	return lexValue
}

func lexValue(lexer *Lexer) lexState {
	for {
		if lexer.isEOF() {
			lexer.pos--
			lexer.putToken(token.TOKEN_VALUE)
			return lexEOF
		}

		if unicode.IsSpace(lexer.next()) {
			lexer.pos--
			lexer.putToken(token.TOKEN_VALUE)
			lexer.skipBlank()

			return lexBegin
		}

		lexer.increment()
	}
}

func lexEOF(lexer *Lexer) lexState {
	lexer.putToken(token.TOKEN_EOF)
	return nil
}
