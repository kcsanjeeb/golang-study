package main

import "fmt"

func main() {
	// Break example
	fmt.Println("Break example:")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("\n----------------------------")

	// Continue example
	fmt.Println("Continue example:")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("\n----------------------------")

	// Labels example
	fmt.Println("Labels example:")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	fmt.Println("\n----------------------------")

	// Goto example
	fmt.Println("Goto example:")
	i := 0
start:
	if i < 5 {
		fmt.Println(i)
		i++
		goto start
	}

	fmt.Println("\n----------------------------")

	// Switch statement
	fmt.Println("Switch statement:")
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("It's Monday")
	case "Tuesday":
		fmt.Println("It's Tuesday")
	default:
		fmt.Println("Other day")
	}
}
