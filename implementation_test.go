package lab2

import (
	"testing"
)

func TestPostfixToLisp(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Простий вираз
		{"4 2 -", "(- 4 2)"},
		{"3 2 +", "(+ 3 2)"},
		{"4 2 * 3 +", "(+ (* 4 2) 3)"},
		{"5 3 2 ^ *", "(* 5 (pow 3 2))"},

		// Складніший вираз
		{"4 2 - 3 2 ^ * 5 +", "(+ (* (- 4 2) (pow 3 2)) 5)"},
		{"3 4 + 5 6 + *", "(* (+ 3 4) (+ 5 6))"},

		// Неправильні дані
		{"", ""}, 
		{"4 2 @", ""}, 
		{"3 + 5", ""}, 
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := postfixToLisp(test.input)
			if result != test.expected {
				t.Errorf("For input %s, expected %s, but got %s", test.input, test.expected, result)
			}
		})
	}
}
