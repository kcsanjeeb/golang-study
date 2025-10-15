package main

import "fmt"

const x = 10       // untyped
const y int32 = 15 // typed

const (
	g               = 10
	h         int32 = 20
	lesson          = "lesson 4"
	isRunning       = false
	character       = 'a'
	isTrue          = 1 > 2
)

func main() {
	var a int
	a = x // can use typed constant
	fmt.Println(a)

	var b float64
	b = x // can use typed constant
	fmt.Println(b)

	var c int
	c = int(y) // Need to typecase , cannot use typed constant
	fmt.Println(c)

	// var d string
	// d = x		we cannot do this because type different

	// RUNTIME CONSTANT
	// const z = a + int(b) // this wont work because variables can be changed  and will not be a constant

	// const z = sum(a,b) // not possible

	const z = len("string") // allow
	fmt.Println(z)

	const k = complex(10.2, 100.9)
	fmt.Println(k)

	const l = imag(k)
	fmt.Println(l)

	fmt.Println(g, h)

	fmt.Println(lesson, isRunning, character, isTrue)
}

func sum(a, b int) int {
	return a + b
}
