package main

import "fmt"

func main() {
	// array declaration basic form
	var a [5]int
	fmt.Println(a)

	// array declaration with elements
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// sparse array declaration
	c := [5]int{1, 3: 44, 5}
	fmt.Println(c)

	//implicit length declaration
	d := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(d)
	// len function to calculate the length of an array
	fmt.Println(len(d))

	// Accessing array elements
	e := d[3]
	fmt.Println(e)

	// TWO DIMENTION ARRAY
	f := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(f)
}
