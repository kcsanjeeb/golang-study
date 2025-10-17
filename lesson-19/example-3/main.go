package main

import "fmt"

func main() {
	fmt.Println("=== Starting Database Transaction ===")

	// Simulate multiple cleanup operations
	defer fmt.Println("5. Transaction COMMITTED")      // Last in - First out (5th)
	defer fmt.Println("4. Releasing connection")       // Fourth in - Second out
	defer fmt.Println("3. Closing temporary files")    // Third in - Third out
	defer fmt.Println("2. Unlocking user records")     // Second in - Fourth out
	defer fmt.Println("1. Logging operation complete") // First in - Last out (1st)

	fmt.Println("Working on database operations...")
	fmt.Println("About to finish function...")

	// All deferred functions execute HERE in REVERSE order
}
