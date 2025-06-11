package lexer

import "toylang/token"

type Lexer struct {
	input         string
	position      int  // current char
	real_position int  // next char
	ch            byte // char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.read_char()
	return l
}

// It only handles ASCII
func (l *Lexer) read_char() {
	if l.real_position >= len(l.input) {
		l.ch = 0 // ASCII code for NUL character
	} else {
		l.ch = l.input[l.real_position]
	}
	l.position = l.real_position
	l.real_position += 1
}
func new_token(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = new_token(token.ASSIGN, l.ch)
	case ';':
		tok = new_token(token.SEMICOLON, l.ch)
	case '(':
		tok = new_token(token.LPAREN, l.ch)
	case ')':
		tok = new_token(token.RPAREN, l.ch)
	case '{':
		tok = new_token(token.LBRACE, l.ch)
	case '}':
		tok = new_token(token.RBRACE, l.ch)
	case ',':
		tok = new_token(token.COMMA, l.ch)
	case '+':
		tok = new_token(token.PLUS, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.read_char()
	return tok
}
