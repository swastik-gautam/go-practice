package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrInvalidAge = errors.New("age must be between 0 and 120")

func checkAge(age int) error {
	if age < 0 || age > 120 {
		return ErrInvalidAge
	}

	return nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		// %w wraps the original error
		return "", fmt.Errorf("failed to read file %q: %w", filename, err)
	}

	return string(data), nil
}

func main() {

	// 1. Handling a normal custom error
	err := checkAge(150)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Valid age")
	}

	// 2. Handling errors from a function returning multiple values
	result, err := divide(10, 0)

	if err != nil {
		fmt.Println("Division error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// 3. Wrapped errors
	content, err := readFile("hello.txt")

	if err != nil {

		fmt.Println("Error:", err)

		// Check whether the wrapped error is because
		// the file doesn't exist
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("The file does not exist")
		}

	} else {
		fmt.Println("File content:")
		fmt.Println(content)
	}
}
