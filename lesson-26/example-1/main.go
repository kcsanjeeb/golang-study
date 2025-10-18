package main

import "fmt"

func panicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()
	panic("Something went Wrong !")
}

func main() {
	fmt.Println("Program has started")
	panicExample()
	fmt.Println("Exited !!")

}
