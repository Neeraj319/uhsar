package main

import "fmt"

type TokenType int

const (
	// single-character token
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two characters token
	BANG
	BANG_EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals
	IDENTIFIER
	STRING
	NUMBER

	// keywords
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
	VAR
	WHILE

	EOF
)

type Token struct {
	Type TokenType
	Lexeme    string
	Lieteral  any
	Line      int
}

func ToString(token *TokenType) string {
	return_string := fmt.Sprintf("%d", )
}
