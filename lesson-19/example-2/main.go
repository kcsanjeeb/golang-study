package main

import "fmt"

// Example demonstrating the three execution conditions
func demonstrateDeferConditions() {
	fmt.Println("\n=== Condition 1: Normal Return ===")
	defer fmt.Println("Deferred: Normal return")
	fmt.Println("Function executing...")
	return // Deferred function executes here
}

func demonstrateDeferCondition2() {
	fmt.Println("\n=== Condition 2: End of Function Body ===")
	defer fmt.Println("Deferred: End of function body")
	fmt.Println("Function executing...")
	// No return statement - deferred function executes at closing brace
}

func demonstrateDeferCondition3() {
	fmt.Println("\n=== Condition 3: Function Panics ===")
	defer fmt.Println("Deferred: After panic")
	fmt.Println("About to panic...")
	panic("Something went wrong!")
	// Deferred function executes before panic propagates
}

func demonstrateArgumentEvaluation() {
	fmt.Println("\n=== Argument Evaluation Timing ===")
	x := "initial"
	defer fmt.Println("Deferred with x:", x) // x is captured as "initial"
	x = "changed"
	fmt.Println("x after change:", x)
}

func demonstrateLIFOOrder() {
	fmt.Println("\n=== LIFO Execution Order ===")
	defer fmt.Println("Deferred 1: First")
	defer fmt.Println("Deferred 2: Second")
	defer fmt.Println("Deferred 3: Third")
	fmt.Println("Main execution")
	// Executes: Third, Second, First (LIFO)
}

func demonstrateReturnTiming() (result int) {
	fmt.Println("\n=== Return Value Timing ===")
	defer func() {
		fmt.Println("Deferred: result is", result)
		result++ // Can modify named return values
	}()
	result = 42
	fmt.Println("Before return: result is", result)
	return result // Deferred function executes after return value is set
}

func main() {
	// Run all demonstrations
	demonstrateDeferConditions()
	demonstrateDeferCondition2()

	// Uncomment to see panic condition (will stop execution)
	// demonstrateDeferCondition3()

	demonstrateArgumentEvaluation()
	demonstrateLIFOOrder()

	result := demonstrateReturnTiming()
	fmt.Println("Final result:", result)

	// Practical example: file cleanup
	fmt.Println("\n=== Practical Example: Resource Cleanup ===")
	simulateFileOperation()
}

func simulateFileOperation() {
	fmt.Println("Opening file...")
	defer fmt.Println("Deferred: Closing file") // Always executes
	defer fmt.Println("Deferred: Flushing buffers")

	fmt.Println("Writing data to file...")
	if someCondition := true; someCondition {
		fmt.Println("Operation successful")
		return // Deferred functions still execute
	}
	fmt.Println("This won't print due to return above")
}
