package main

import "fmt"

func main() {
	s := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 190)
	fmt.Println(s)
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
