package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	// Calculate time two hours from now
	twoHours := 2 * time.Hour
	futureTime := now.Add(twoHours)
	fmt.Println("Time in two hours from now ", futureTime)

	// Measure the execution time of a code block
	start := time.Now()
	a := 1
	for i := 0; i < 10000000; i++ {
		a++
	}
	elapsed := time.Since(start)
	fmt.Println("Loop Executed Time:", elapsed)
}
