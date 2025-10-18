package main

import (
	"errors"
	"fmt"
)

func firstFunction() error {
	return fmt.Errorf("original error: something went wrong in first function.")
}

func secondFunction() error {
	firstErr := firstFunction()
	if firstErr != nil {
		secondErr := errors.New("failed in second function ")
		return errors.Join(firstErr, secondErr)
		// return fmt.Errorf("Failed in second function : %w", firstErr)
	}
	return nil
}

func main() {
	err := secondFunction()
	fmt.Println(err)
}
