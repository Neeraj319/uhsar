package main

// import (
// 	"fmt"
// )

type Parser struct {
	Tokens  []Token
	Current int
}

func CreateParser(tokens []Token) *Parser {
	return &Parser{Tokens: tokens, Current: 0}
}

func (p *Parser) Parse() Expr {
	return p.experssion()
}

func (p *Parser) experssion() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparision()
	for p.match(BANG_EQUAL, EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparision()
		expr = CreateBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) comparision() Expr {
	expr := p.term()
	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = CreateBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()
	for p.match(MINUS, PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = CreateBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()
	for p.match(SLASH, STAR) {
		operator := p.previous()
		right := p.unary()
		expr = CreateBinary(expr, operator, right)
	}
	return expr
}

func (p *Parser) unary() Expr {
	if p.match(BANG, MINUS) {
		operator := p.previous()
		right := p.unary()
		return CreateUnary(operator, right)
	}
	return p.primary()
}
func (p *Parser) primary() Expr {
	if p.match(FALSE) {
		return CreateLiteral(false)
	}
	if p.match(TRUE) {
		return CreateLiteral(true)
	}
	if p.match(NIL) {
		return CreateLiteral(nil)
	}
	if p.match(NUMBER, STRING) {
		return CreateLiteral(p.previous().Literal)
	}
	if p.match(LEFT_PAREN) {
		expr := p.experssion()
		p.consume(RIGHT_PAREN, "Expect ')' after expression.")
		return CreateGrouping(expr)
	}
	panic(ParsingError(p.peek(), "Expected expression"))
}

func (p *Parser) consume(tokenType TokenType, message string) Token {
	if p.check(tokenType) {
		return p.advance()
	}
	panic(ParsingError(p.peek(), message))
}

func (p *Parser) match(tokenTypes ...TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false

}

func (p *Parser) check(tokenType TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.Current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == EOF
}

func (p *Parser) peek() Token {
	return p.Tokens[p.Current]
}

func (p *Parser) previous() Token {
	return p.Tokens[p.Current-1]
}
