package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Message string
	Wrapped error
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Wrapped)
}

func (e CustomError) Unwrap() error {
	return e.Wrapped
}

func SomeFunction() error {
	return CustomError{
		Message: "original error: something went wrong",
		Wrapped: errors.New("wrapped error"),
	}
}

func main() {
	err := SomeFunction()
	fmt.Println("Error :", err)
	err = errors.Unwrap(err)
	innerError := fmt.Errorf("innermost: %w", err)
	fmt.Println("Innermost Error :", innerError)
}
