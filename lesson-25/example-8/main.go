package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Message string
	Status  int
}

func (e CustomError) Error() string {
	return e.Message
}

func SomeFunction() error {
	return CustomError{Message: "original error: something went wrong", Status: 400}
}

func main() {
	err := SomeFunction()
	var customErr CustomError
	if errors.As(err, &customErr) {
		fmt.Println("Extracted CustomError: ", customErr)
	}
}
