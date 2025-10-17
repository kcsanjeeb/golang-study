package main

import "fmt"

func main() {
	var emptyInterface any
	emptyInterface = "Hello World"
	if str, ok := emptyInterface.(string); ok {
		fmt.Println("The underlying value is a string", str)
	}
}
