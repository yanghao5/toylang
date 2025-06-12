package main

import (
	"fmt"
	"toylang/lexer"
)

func main() {
	input := `let five = 5;
let ten = 10;
`
	l := lexer.New(input)

	for tok := l.NextToken(); tok.Literal != ""; tok = l.NextToken() {
		fmt.Println(tok.Literal, "||", tok.Type)
	}
}
