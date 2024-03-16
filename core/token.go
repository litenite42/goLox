package core

import (
	"fmt"
)

type Token struct {
	ttype   TokenType
	lexeme  string
	literal interface{}
	line    int
}

func newToken(ttype TokenType, lexeme string, literal interface{}, line int) *Token {
	return &Token{ttype, lexeme, literal, line}
}

func (t Token) String() string {
	return fmt.Sprintf("%d %s %s", t.ttype, t.lexeme, t.literal)
}
