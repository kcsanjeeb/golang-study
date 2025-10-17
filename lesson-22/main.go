package main

import "fmt"

type Age int

type Human struct {
	Age int
}

type Student Human

func sum(a Age) {
	fmt.Printf("Age %d ", a)
}

func main() {
	var young Age = 10
	sum(young)
}
