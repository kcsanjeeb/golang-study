package main

import "fmt"

func main() {
	s := []int{1, 2}
	fmt.Println(s) // [1 2]
	somethingslice(s)
	fmt.Println(s) // [100 2] when changing values , [1 2] when appending values

	fmt.Println("in original function", len(s))
}

func somethingslice(s []int) {
	// s[0] = 100		// changing values
	s = append(s, 1000) // appending values
	// Inside the function slice copy was created, when i do append , its appending beyong the lenght of that slice
	// the slice length is updated inside the function
	// but since it is a copy the original slice length is not incremented
	fmt.Println("in copy function: ", len(s))

}
