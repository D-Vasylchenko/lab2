package lab2

import (
	"fmt"
	"strings"
	"unicode"
)

func isOperand(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func postfixToLisp(postfix string) string {
	var stack []string

	tokens := strings.Fields(postfix)

	for _, token := range tokens {
		if isOperand(token) {
			stack = append(stack, token)
		} else {
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if token == "^" {
				token = "pow"
			}

			stack = append(stack, fmt.Sprintf("(%s %s %s)", token, operand1, operand2))
		}
	}

	return stack[0]
}