package main

import "fmt"

type Calculator struct {
}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func AddFunction(a, b int) int {
	return a + b
}

func ArthmeticOperation(f func(int, int) int, a, b int) int {
	return f(a, b)
}

func main() {
	c := Calculator{}

	fmt.Println(AddFunction(10, 50))

	fmt.Println(ArthmeticOperation(AddFunction, 10, 50))

	fmt.Println(ArthmeticOperation(c.Add, 10, 50))

}
