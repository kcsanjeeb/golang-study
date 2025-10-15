package main

import "fmt"

func main() {
	age := 12

	// Basic if-else
	if age >= 18 {
		fmt.Println("You are an adult!")
	} else {
		fmt.Println("You are not an adult")
	}

	// if-else if ladder
	if age >= 18 {
		fmt.Println("You are an adult!")
	} else if age >= 13 {
		fmt.Println("You are a teenager!")
	} else {
		fmt.Println("You are a child!")
	}

	// Scoped variable in if statement
	if even := isEven(age); even {
		fmt.Println("Age is even!")
	}
}

func isEven(n int) bool {
	return n&1 == 0
}
