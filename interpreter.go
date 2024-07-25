package main

type Interpreter struct{}

func (i *Interpreter) VisitLiteralExpr(literalExpr *Literal) interface{} {
	return literalExpr.Value
}

func (i *Interpreter) VisitGroupingExpr(groupingExpr *Grouping) interface{} {
	return i.evaluate(groupingExpr.Expression)

}

func (i *Interpreter) VisitUnaryExpr(unaryExpr *Unary) interface{} {
	right := i.evaluate(unaryExpr.Right)
	switch unaryExpr.Operator.Type {
	case BANG:
		return !IsTruthy(right)
	case MINUS:
		return -right.(float64)
	}
	return nil
}

func (i *Interpreter) VisitBinaryExpr(binaryExpr *Binary) interface{} {
	left := i.evaluate(binaryExpr.Left)
	right := i.evaluate(binaryExpr.Right)

	switch binaryExpr.Operator.Type {
	case MINUS:
		return left.(float64) - right.(float64)
	case SLASH:
		return left.(float64) / right.(float64)
	case STAR:
		return left.(float64) * right.(float64)
	case PLUS:
		leftDouble, leftOk := left.(float64)
		rightDouble, rightOk := right.(float64)
		if leftOk && rightOk {
			return leftDouble + rightDouble
		}
		leftString, leftOk := left.(string)
		rightString, rightOk := right.(string)
		if leftOk && rightOk {
			return leftString + rightString
		}
		break
	case GREATER:
		return left.(float64) > right.(float64)
	case GREATER_EQUAL:
		return left.(float64) >= right.(float64)
	case LESS:
		return left.(float64) >= right.(float64)
	case LESS_EQUAL:
		return left.(float64) >= right.(float64)
	case BANG_EQUAL:
		return !IsEqual(left, right)
	case EQUAL_EQUAL:
		return IsEqual(left, right)

	}
	return nil

}

func (i *Interpreter) evaluate(expr Expr) interface{} {
	return expr.Accept(i)
}

func Interpret(expr Expr) interface{} {
	interpreter := Interpreter{}
	return interpreter.evaluate(expr)
}

func (ap *Interpreter) VisitRashuExpr(rashuExpr *RashuExpr) interface{} {
	return ""
}
