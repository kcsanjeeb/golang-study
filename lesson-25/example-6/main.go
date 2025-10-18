package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Message string
}

var ErrFirstError = CustomError{Message: "first error"}

func (c CustomError) Error() string {
	return c.Message
}

func SomeFunction() error {
	return ErrFirstError
}

func main() {
	err := SomeFunction()
	fmt.Println(err)

	if err != nil {
		if errors.Is(err, ErrFirstError) {
			fmt.Println("sentinal error found")
		}
	}
}
