package main

import "fmt"

// Basic functions
func double(x int) int {
	return x * 2
}

func triple(x int) int {
	return x * 3
}

// Higher-order function: accepts function as parameter
func apply(f func(int) int, x int) int {
	return f(x)
}

// Function type definition
type operation func(int, int) int

// Function returning a function
func arithmeticOperation(op string) operation {
	switch op {
	case "add":
		return func(a, b int) int {
			return a + b
		}
	case "sub":
		return func(a, b int) int {
			return a - b
		}
	}
	return func(a, b int) int {
		return 0
	}
}

// Function type and factory function
type multiplier func(int) int

func multiplyBy(m int) multiplier {
	return func(i int) int {
		return i * m
	}
}

func main() {
	// Passing functions as arguments
	result := apply(double, 5)
	result2 := apply(triple, 5)

	// Anonymous function assigned to variable
	fiveTime := func(x int) int {
		return 5 * x
	}
	result3 := apply(fiveTime, 5)

	fmt.Println("Double of 5:", result)
	fmt.Println("Triple of 5:", result2)
	fmt.Println("Five times 5:", result3)

	fmt.Println("\n-----------------------")

	// Function returning functions
	var add operation
	add = arithmeticOperation("add")
	result4 := add(1, 1)

	sub := arithmeticOperation("sub")
	result5 := sub(5, 1)

	fmt.Println("1 + 1 =", result4)
	fmt.Println("5 - 1 =", result5)

	fmt.Println("\n-----------------------")

	// Factory functions creating multipliers
	multiplyTwo := multiplyBy(2)
	multiplyThree := multiplyBy(3)
	multiplyFour := multiplyBy(4)

	result6 := multiplyTwo(5)
	result7 := multiplyThree(5)
	result8 := multiplyFour(5)
	fmt.Println("2×5 =", result6, "3×5 =", result7, "4×5 =", result8)
}
