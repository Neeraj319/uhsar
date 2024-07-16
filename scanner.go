package main

import (
	"fmt"
	"strconv"
)

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	line    int
	current int
}

var KEYWORDS = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

func CreateScanner(source string) *Scanner {
	var tokens []Token
	return &Scanner{source: source, tokens: tokens}

}

func (s *Scanner) isAtEnd() bool {
	return s.line >= len(s.source)
}

func (s *Scanner) scanToken() {
	character := s.advance()
	switch character {
	case '(':
		s.addToken(LEFT_PAREN)
		break
	case ')':
		s.addToken(RIGHT_PAREN)
		break

	case '{':
		s.addToken(LEFT_BRACE)
		break
	case '}':
		s.addToken(RIGHT_BRRACE)
		break
	case ',':
		s.addToken(COMMA)
		break
	case '.':
		s.addToken(DOT)
		break
	case '-':
		s.addToken(MINUS)
		break
	case '+':
		s.addToken(PLUS)
		break
	case ';':
		s.addToken(SEMICOLON)
		break
	case '*':
		s.addToken(STAR)
		break
	case '!':
		token := s.match('=')
		if token {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		token := s.match('=')
		if token {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '<':
		token := s.match('=')
		if token {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		token := s.match('=')
		if token {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		s.line++
		break
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case '"':
		s.addString()
		break
	default:
		if IsDigit(character) {
			s.addNumber()
		} else if IsAlpha(character) {
			s.addIdentifier()

		} else {
			fmt.Println("The current values are", s.start, s.current)
			error_string := fmt.Sprintf("Unexpected character %v", character)
			Error(s.line, error_string)
			break
		}
	}

}

func (s *Scanner) addIdentifier() {
	for IsAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	tokenType := KEYWORDS[text]
	if tokenType == "" {
		s.addToken(IDENTIFIER)
	} else {
		s.addToken(tokenType)
	}
}

func (s *Scanner) addNumber() {
	for IsDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && IsDigit(s.peekNext()) {
		s.advance()
	}
	for IsDigit(s.peek()) {
		s.advance()
	}
	value, _ := strconv.ParseFloat(s.source[s.start:s.current], 64)
	s.addTokenWithLiteral(NUMBER, value)
}

func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return '\n'
	}
	return rune(s.source[s.current+1])
}

func (s *Scanner) addString() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		Error(s.line, "Un-terminated string")
		return
	}
	s.advance()
	value := s.source[s.start+1 : s.current-1]
	s.addTokenWithLiteral(STRING, value)
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return 0
	}
	return rune(s.source[s.current])

}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}
	character := rune(s.source[s.current])
	if character != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) advance() rune {
	// if s.current >= len(s.source) {
	// 	return 0
	// }
	// char := rune(s.source[s.current])
	// s.current++
	// return char
	char := s.source[s.current]
	s.current++
	return rune(char)

}

func (s *Scanner) addToken(tokenType TokenType) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{Type: tokenType, Lexeme: text, Literal: nil, Line: s.line})

}

func (s *Scanner) addTokenWithLiteral(tokenType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{Type: tokenType, Lexeme: text, Literal: literal, Line: s.line})
}

func (s *Scanner) ScanTokens() []Token {

	for {
		if s.isAtEnd() {
			token := Token{Type: EOF, Lexeme: "", Literal: nil, Line: s.line}
			s.tokens = append(s.tokens, token)
			s.start = s.current
			return s.tokens
		}
		s.scanToken()
	}

}
