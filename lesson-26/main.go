package main

import "fmt"

func function1() {
	// Registers a deferred func that will print: “Function 1 deffered function called.”
	defer func() {
		fmt.Println("Function 1 deffered function called.")
	}()
	// Calls function2.
	function2()
}

func function2() {
	// Registers a deferred func that will print: “Function 2 deffered function called.”
	defer func() {
		fmt.Println("Function 2 deffered function called.")
	}()
	// Calls panic("Function 2 panicked !").
	panic("Function 2 panicked !")
}

func main() {
	// Registers a deferred func that will print: “Function Main deffered function called.”
	defer func() {
		fmt.Println("Function Main deffered function called.")
	}()

	//Calls function1.
	function1()
}
