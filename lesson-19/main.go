package main

import "fmt"

func main() {
	message := "Hello "

	// Deferred function 1: anonymous function
	defer func(name string) {
		fmt.Println("Deferred 1:", message+name)
	}("Max")

	// Deferred function 2: function variable
	greetingFn := func(name string) {
		fmt.Println("Deferred 2:", message+name)
	}
	defer greetingFn("Alice")
	defer greetingFn("Bob")

	// Change message after defer statements
	message = "Hi "

	fmt.Println("Main function executing...")

	// Deferred functions will execute here when main() ends
	// They execute in LIFO order: Bob, Alice, Max
}
