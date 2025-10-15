package main

import "fmt"

func main() {
	// For with condition
	a := 10
	for a > 0 {
		fmt.Println(a)
		a--
	}

	fmt.Printf("\n----------------------------\n")

	// Infinite loop with break
	b := 10
	for {
		fmt.Println(b)
		b--
		if b == 0 {
			break
		}
	}

	fmt.Printf("\n----------------------------\n")

	// For with clause
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n----------------------------\n")

	// For with range clause - slices
	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range c {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	// For with range clause - strings
	d := "你好"
	for index, runeValue := range d {
		fmt.Printf("index: %d, rune value: %d\n", index, runeValue)
	}
}
