package main

import (
	"fmt"

	_ "github.com/kcsanjeeb/golang-study/lesson-29/init"
)

func init() {
	fmt.Println("Init 1 called")
}
func init() {
	fmt.Println("Init 2 called")
}
func init() {
	fmt.Println("Init 4 called")
}
func init() {
	fmt.Println("Init 3 called")
}

func main() {
	fmt.Println("Main function called")
}
