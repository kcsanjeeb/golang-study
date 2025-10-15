package main

import "fmt"

var myInteger1 int = 10

func main() {
	var a, b, c string
	a = "Hello"
	b = "World"
	c = "!!"
	fmt.Println(a + " " + b + " " + c)

	fmt.Println(myInteger1)
	something()

	// Explicit type assignment
	var myInt int = 10
	var myString string = "hello world !"
	fmt.Println(myInt)
	fmt.Println(myString)

	// Implicit type assignment
	var age = 25
	fmt.Println(age)

	// Multiple variable declaration
	var year, month = 2025, "Jan"
	fmt.Println(year, month)

	postcode := 55800     // var postcode int = 55800
	address := "shenzhen" // var address string = "shenzhen"
	fmt.Println(postcode, address)

	otp_code, language := 1234, "Chinese"
	// var otp_code int = 1234
	// var language string = "Chinese"

	fmt.Println(language, otp_code)
}

func something() {
	fmt.Println("Something: ", myInteger1)
}
