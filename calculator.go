package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Function to calculate the value of an expression
func Calc(expression string) (float64, error) {
	tokens := tokenize(expression)
	if len(tokens) == 0 {
		return 0, errors.New("empty expression")
	}
	return evaluate(tokens)
}

// Tokenizer to split the input expression into tokens
func tokenize(expression string) []string {
	var tokens []string
	var current strings.Builder

	for _, ch := range expression {
		if unicode.IsSpace(ch) {
			continue
		}
		if isOperator(ch) || ch == '(' || ch == ')' {
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			tokens = append(tokens, string(ch))
		} else if unicode.IsDigit(ch) || ch == '.' {
			current.WriteRune(ch)
		} else {
			return nil // Invalid character found
		}
	}

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}
	return tokens
}

// Function to determine if a character is an operator
func isOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

// Evaluate the expression using the Shunting Yard algorithm
func evaluate(tokens []string) (float64, error) {
	// Implement filement logic later
	// For simplicity, we will consider only the addition and subtraction here.
	output := []string{}
	operatorStack := []string{}

	for _, token := range tokens {
		if isNumber(token) {
			output = append(output, token)
		} else if isOperator(rune(token[0])) {
			for len(operatorStack) > 0 && precedence(operatorStack[len(operatorStack)-1]) >= precedence(token) {
				output = append(output, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			operatorStack = append(operatorStack, token)
		} else if token == "(" {
			operatorStack = append(operatorStack, token)
		} else if token == ")" {
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != "(" {
				output = append(output, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			if len(operatorStack) == 0 {
				return 0, errors.New("mismatched parentheses")
			}
			operatorStack = operatorStack[:len(operatorStack)-1] // pop the "("
		}
	}

	for len(operatorStack) > 0 {
		output = append(output, operatorStack[len(operatorStack)-1])
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	return calculateRPN(output)
}

// Function to check if a string is a number
func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// Define precedence of operators
func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

// Evaluate RPN expression
func calculateRPN(tokens []string) (float64, error) {
	stack := []float64{}

	for _, token := range tokens {
		if isNumber(token) {
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		} else if isOperator(rune(token[0])) {
			if len(stack) < 2 {
				return 0, errors.New("invalid expression")
			}
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				result = a / b
			default:
				return 0, errors.New("unsupported operator")
			}
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}
	return stack[0], nil
}

func main() {
	exampleExpr := "3 + 5 * ( 2 - 8 )"
	result, err := Calc(exampleExpr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
