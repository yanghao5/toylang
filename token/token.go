package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// end of file
	EOF = "EOF"

	// unknown
	ILLEGAL = "ILLEGAL"

	// Identifier
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	// Operator
	ASSIGN = "="
	PLUS   = "+"

	// Delimiter
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// keyword
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
