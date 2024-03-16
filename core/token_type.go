package core

type TokenType int

const (
	LPAREN TokenType = iota
	RPAREN
	LBRACE
	RBRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	//
	BANG
	BANG_EQ
	EQ
	EQ_EQ
	GTR
	GTR_EQ
	LS
	LS_EQ
	//
	IDENT
	STRING
	NUMBER
	//
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	WHILE
	//
	EOF
)
