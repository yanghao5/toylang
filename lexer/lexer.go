package lexer

import (
	"toylang/token"
)

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

func new_token(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// _ is considered letter
func is_letter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
func is_digit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skip_whitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		//fmt.Println("Skipping whitespace: ", string(l.ch), "position:", l.position, "real_position:", l.real_position) // Debug output
		l.read_char()
	}
}

// It only handles ASCII
func (l *Lexer) read_char() {
	if l.real_position >= len(l.input) {
		l.ch = 0 // ASCII code for NUL character
	} else {
		l.ch = l.input[l.real_position]
	}
	//fmt.Println("l.ch: ", string(l.ch), "position:", l.position, "real_position:", l.real_position) // Debug output
	l.position = l.real_position
	l.real_position += 1
}

func (l *Lexer) read_identifier() string {
	position := l.position
	for is_letter(l.ch) {
		l.read_char()
	}
	return l.input[position:l.position]
}

func (l *Lexer) read_number() string {
	position := l.position
	for is_digit(l.ch) {
		l.read_char()
	}
	return l.input[position:l.position]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skip_whitespace()
	//fmt.Println("l.ch ", string(l.ch))

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
	default:
		if is_letter(l.ch) {
			tok.Literal = l.read_identifier() // Need to exit early, because read_identifier() called l.read_char()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if is_digit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.read_number()
			return tok
		} else {
			tok = new_token(token.ILLEGAL, l.ch)
		}
	}
	l.read_char()
	return tok
}
