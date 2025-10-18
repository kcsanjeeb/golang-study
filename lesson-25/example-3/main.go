package main

import "fmt"

func firstFunction() error {
	return fmt.Errorf("original error: something went wrong in first function.")
}

func secondFunction() error {
	firstErr := firstFunction()
	if firstErr != nil {
		return fmt.Errorf("Failed in second function : %w", firstErr)
	}
	return nil
}

func main() {
	err := secondFunction()
	fmt.Println(err)
}
