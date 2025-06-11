package lab2

import (
	"bufio"
	"fmt"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (h *ComputeHandler) Compute() error {
	// Читаємо вхідні дані
	scanner := bufio.NewScanner(h.Input)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read input: %v", err)
	}

	result := postfixToLisp(input)
	if result == "" {
		return fmt.Errorf("syntax error in input expression")
	}

	_, err := fmt.Fprintln(h.Output, result)
	return err
}
