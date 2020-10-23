package lexer_no_chan

import "github.com/doughnuty/GOparser/token"

func (lexer *Lexer) changeState(state lexState) {
	lexer.state = state
}

// return next token from channel
func (lexer *Lexer) NextToken() token.Token {
	// if we reached the end no reason to search for new token
	if lexer.Current.Mod == token.TOKEN_EOF {
		return lexer.Current
	}

	return lexer.state(lexer)
}

// put token into token channel
func (lexer *Lexer) putToken(tokenMod token.TokenMod) token.Token {
	t := token.Token{Mod: tokenMod, Value: lexer.input[lexer.start:lexer.pos]}
	lexer.start = lexer.pos
	return t
}
