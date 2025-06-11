package main

import (
	"flag"
	"fmt"
	"io"
	"lab2"
	"os"
	"strings"
)

func main() {
	// Оголошуємо флаги для опцій
	expressionFlag := flag.String("e", "", "Postfix expression")
	fileFlag := flag.String("f", "", "File containing postfix expression")
	outputFlag := flag.String("o", "", "File to write the result")

	// Парсимо флаги
	flag.Parse()

	// Перевірка, чи вказано обидва параметри -e та -f
	if *expressionFlag != "" && *fileFlag != "" {
		fmt.Fprintln(os.Stderr, "Error: cannot specify both -e and -f options.")
		os.Exit(1)
	}

	// Визначаємо джерело для вводу
	var input io.Reader
	if *expressionFlag != "" {
		// Якщо вказано -e, беремо вираз з командного рядка
		input = strings.NewReader(*expressionFlag)
	} else if *fileFlag != "" {
		// Якщо вказано -f, відкриваємо файл для читання
		file, err := os.Open(*fileFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		// Якщо не вказано ані -e, ані -f, виводимо помилку
		fmt.Fprintln(os.Stderr, "Error: must specify either -e or -f.")
		os.Exit(1)
	}

	// Визначаємо, куди записувати результат
	var output io.Writer = os.Stdout
	if *outputFlag != "" {
		// Якщо вказано -o, записуємо у файл
		file, err := os.Create(*outputFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating output file:", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	}

	// Створюємо екземпляр ComputeHandler та викликаємо Compute()
	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	// Обчислюємо результат
	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
