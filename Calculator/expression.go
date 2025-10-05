package main

type Expression struct {
	leftExpression  ArthimeticExpression
	rightExpression ArthimeticExpression
	operator        string
}

func createExpression(leftExpression ArthimeticExpression, rightExpression ArthimeticExpression, operator string) *Expression {
	return &Expression{
		leftExpression:  leftExpression,
		rightExpression: rightExpression,
		operator:        operator,
	}
}

func (e *Expression) Evaluate() int {
	switch e.operator {
	case "ADD":
		return e.leftExpression.Evaluate() + e.rightExpression.Evaluate()
	case "SUB":
		return e.leftExpression.Evaluate() - e.rightExpression.Evaluate()
	case "MUL":
		return e.leftExpression.Evaluate() * e.rightExpression.Evaluate()
	case "DIV":
		return e.leftExpression.Evaluate() / e.rightExpression.Evaluate()
	default:
		return 0
	}
}
