package main

func main() {
	num1 := &Number{value: 10}
	num2 := &Number{value: 5}
	expression := createExpression(num1, num2, "ADD")
	num3 := &Number{value: 2}
	expression = createExpression(expression, num3, "MUL")
	result := expression.Evaluate()
	println("Result:", result) // Output: Result: 30
}
