package main

import (
	"fmt"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero %f", b)
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
