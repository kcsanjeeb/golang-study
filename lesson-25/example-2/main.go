package main

import (
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("Error %d:  %s ", c.Code, c.Message)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, CustomError{Code: 1, Message: "Cannot divide by 0"}
	}
	return a / b, nil
}

func main() {
	result, err := Divide(10, 0)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("result", result)

}
