package main

import (
	"fmt"
)

type TokenType string

const (
	// single-character token
	LEFT_PAREN   TokenType = "LEFT_PAREN"
	RIGHT_PAREN            = "RIGHT_PAREN"
	LEFT_BRACE             = "LEFT_BRACE"
	RIGHT_BRRACE           = "RIGHT_BRRACE"
	COMMA                  = "COMMA"
	DOT                    = "DOT"
	MINUS                  = "MINUS"
	PLUS                   = "PLUS"
	SEMICOLON              = "SEMICOLON"
	SLASH                  = "SLASH"
	STAR                   = "STAR"
	EQUAL                  = "EQUAL"

	// One or two characters token
	BANG          = "BANG"
	BANG_EQUAL    = "BANG_EQUAL"
	EQUAL_EQUAL   = "EQUAL_EQUAL"
	GREATER       = "GREATER"
	GREATER_EQUAL = "GREATER_EQUAL"
	LESS          = "LESS"
	LESS_EQUAL    = "LESS_EQUAL"

	// Literals
	IDENTIFIER = "IDENTIFIER"
	STRING     = "STRING"
	NUMBER     = "NUMBER"

	// keywords
	AND    = "AND"
	CLASS  = "CLASS"
	ELSE   = "ELSE"
	FALSE  = "FALSE"
	FUN    = "FUN"
	FOR    = "FOR"
	IF     = "IF"
	NIL    = "NIL"
	OR     = "OR"
	PRINT  = "PRINT"
	RETURN = "RETURN"
	SUPER  = "SUPER"
	THIS   = "THIS"
	TRUE   = "TRUE"
	VAR    = "VAR"
	WHILE  = "WHILE"

	EOF = "EOF"
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func (t *Token) ToString() string {
	return_string := fmt.Sprintf("type: %s, lexeme value: %s, literal value: %v", t.Type, t.Lexeme, t.Literal)
	return return_string
}

func CreateToken(tokenType TokenType, lexeme string, literal interface{}, line int) *Token {
	return &Token{Type: tokenType, Lexeme: lexeme, Literal: literal, Line: line}

}
