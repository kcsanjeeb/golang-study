package main

import "fmt"

var (
	a = 10 // Package-level variable
)

func main() {
	fmt.Println(a) // 10 (package-level)
	{
		a := 15        // Inner block variable - SHADOWS package-level 'a'
		fmt.Println(a) // 15 (inner block)
	}
	fmt.Println(a) // 10 (package-level again - shadow ends)
}
