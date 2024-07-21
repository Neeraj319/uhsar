package main

type VisitorInterface interface {
	VisitBinaryExpr(expr *Binary) interface{}
	VisitGroupingExpr(expr *Grouping) interface{}
	VisitLiteralExpr(expr *Literal) interface{}
	VisitUnaryExpr(expr *Unary) interface{}
}

type Expr interface {
	Accept(visitor VisitorInterface) interface{}
}

type Binary struct {
	Left     Expr
	Operator Token
	Right    Expr
}

type Grouping struct {
	Expression Expr
}

type Literal struct {
	Value interface{}
}

type Unary struct {
	Operator Token
	Right    Expr
}

func CreateBinary(left Expr, operator Token, right Expr) *Binary {
	return &Binary{Left: left, Operator: operator, Right: right}
}

func CreateGrouping(expression Expr) *Grouping {
	return &Grouping{Expression: expression}
}

func CreateLiteral(value interface{}) *Literal {
	return &Literal{Value: value}
}

func CreateUnary(operator Token, right Expr) *Unary {
	return &Unary{Operator: operator, Right: right}
}

func (b *Binary) Accept(visitor VisitorInterface) interface{} {
	return visitor.VisitBinaryExpr(b)
}

func (g *Grouping) Accept(visitor VisitorInterface) interface{} {
	return visitor.VisitGroupingExpr(g)
}

func (l *Literal) Accept(visitor VisitorInterface) interface{} {
	return visitor.VisitLiteralExpr(l)
}

func (u *Unary) Accept(visitor VisitorInterface) interface{} {
	return visitor.VisitUnaryExpr(u)
}

