package main

import "fmt"

// Closure example: counter factory function
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Higher-order function using closures
func sliceOperations(s []int, op func(int) int) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = op(v)
	}
	return result
}

func main() {
	// Using closure to maintain state
	increment := counter()
	fmt.Println("Counter:", increment())
	fmt.Println("Counter:", increment())
	fmt.Println("Counter:", increment())

	fmt.Println("\n-----------------------")

	// Immediately invoked anonymous function with parameter
	func(msg string) {
		fmt.Println("Message:", msg)
	}("Hello World!")

	// Immediately invoked anonymous function without parameters
	func() {
		fmt.Println("This is an anonymous function!")
	}()

	fmt.Println("\n-----------------------")

	// Using anonymous function as callback
	numbers := []int{1, 2, 3, 4, 5}

	squared := sliceOperations(numbers, func(i int) int {
		return i * i
	})
	fmt.Println("Squared numbers:", squared)

	// Another closure example: multiplier factory
	multiplier := func(factor int) func(int) int {
		return func(x int) int {
			return x * factor
		}
	}

	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println("Double 5:", double(5))
	fmt.Println("Triple 5:", triple(5))

	fmt.Println("\n-----------------------")

	// Closure with multiple captured variables
	accumulator := func() func(int) int {
		total := 0
		return func(x int) int {
			total += x
			return total
		}
	}()

	fmt.Println("Accumulator:", accumulator(10))
	fmt.Println("Accumulator:", accumulator(5))
	fmt.Println("Accumulator:", accumulator(3))
}
