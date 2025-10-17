package main

import "fmt"

func printValue(value any) {
	fmt.Println(value)
}

func main() {
	mixedSlice := []interface{}{42, "hello", 3.14, true}
	for _, value := range mixedSlice {
		printValue(value)
	}
}
