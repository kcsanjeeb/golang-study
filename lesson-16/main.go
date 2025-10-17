package main

import "fmt"

func main() {
	// Calling variadic function with multiple arguments
	s := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 190)
	fmt.Println("Sum:", s)

	// Calling with no arguments
	emptySum := sum()
	fmt.Println("Empty sum:", emptySum)

	// Calling with slice using spread operator
	numbers := []int{10, 20, 30}
	sliceSum := sum(numbers...)
	fmt.Println("Slice sum:", sliceSum)

	// Mixed parameters example
	greetPeople("Meeting:", "Alice", "Bob", "Charlie")

	// Real-world example: joining strings
	path := joinPath("usr", "local", "bin", "go")
	fmt.Println("Path:", path)
}

// Variadic function that sums any number of integers
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Function with regular and variadic parameters
func greetPeople(prefix string, names ...string) {
	fmt.Print(prefix + " ")
	for _, name := range names {
		fmt.Print(name + " ")
	}
	fmt.Println()
}

// Practical example: building file paths
func joinPath(parts ...string) string {
	path := ""
	for i, part := range parts {
		if i > 0 {
			path += "/"
		}
		path += part
	}
	return path
}
