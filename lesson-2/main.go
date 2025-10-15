package main

import "fmt"

func main() {
	// STRING
	var myString string
	fmt.Println(myString) // prints nothing
	myString = "Welcome to coding."
	fmt.Println(myString) // prints the string

	// INTEGER
	var myInt int
	fmt.Println(myInt) // prints 0
	myInt = 10
	fmt.Println(myInt)

	// BOOLEAN
	// zero value for bool is false
	var myBool bool
	fmt.Println(myBool)
	myBool = true
	fmt.Println(myBool)
}
