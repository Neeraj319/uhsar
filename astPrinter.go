package main

import "fmt"

type AstPrinter struct{}

func AstPrint(expr Expr) string {
	return expr.Accept(&AstPrinter{}).(string)
}

func Parenthesize(name string, exprs ...Expr) string {
	var str string
	str += "(" + name
	for _, expr := range exprs {
		str += " " + AstPrint(expr)
	}
	str += ")"
	return str
}

func (ap *AstPrinter) VisitBinaryExpr(binaryExpr *Binary) interface{} {
	return Parenthesize(binaryExpr.Operator.Lexeme, binaryExpr.Left, binaryExpr.Right)
}

func (ap *AstPrinter) VisitGroupingExpr(groupingExpr *Grouping) interface{} {
	return Parenthesize("group", groupingExpr.Expression)
}

func (ap *AstPrinter) VisitUnaryExpr(unaryExpr *Unary) interface{} {
	return Parenthesize(unaryExpr.Operator.Lexeme, unaryExpr.Right)
}

func (ap *AstPrinter) VisitLiteralExpr(literalExpr *Literal) interface{} {
	if literalExpr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", literalExpr.Value)
}
