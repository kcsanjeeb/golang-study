package main

import "fmt"

type Size int

const (
	ExtraSmall Size = iota * 2
	Small
	Medium
	Large
	ExtraLarge
)

const (
	New int = iota
)

func printSize(s Size) {
	switch s {
	case ExtraSmall:
		fmt.Println("Extra Small")
	case Small:
		fmt.Println("Small")
	case Medium:
		fmt.Println("Medium")
	case Large:
		fmt.Println("Large")
	case ExtraLarge:
		fmt.Println("Extra Large")

	}
}

func main() {
	fmt.Println(ExtraSmall, Small, Medium, Large, ExtraLarge)
}
