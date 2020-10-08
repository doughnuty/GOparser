package errors

const (
	LEXER_ERROR_UNEXPECTED_EOF string = "unexpected end of file"
	LEXER_ERROR_MISSING_COLON  string = "missing a colon after the key word"
	LEXER_BAD_INDENTATION      string = "indentation error: incorrect number of spaces"
	LEXER_MISSING_BRACKET      string = "expected right bracket found new line"
)
