package main

import "fmt"

// Basic function with no parameters and no return value
func sayHi() {
	fmt.Println("Hello!")
}

// Function with parameters
func greetPerson(firstName string, lastName string) {
	fmt.Printf("Hello %s %s!\n", firstName, lastName)
}

// Function with return value
func getFullName(firstName string, lastName string) string {
	return firstName + " " + lastName
}

// Function with multiple return values
func getUserInfo(firstName string, lastName string, age int) (string, int) {
	fullName := firstName + " " + lastName
	return fullName, age
}

// Function with named return values
func getPersonDetails(firstName string, lastName string) (fullName string, nameLength int) {
	fullName = firstName + " " + lastName
	nameLength = len(fullName)
	return // naked return - uses named return values
}

func main() {
	// Calling basic function
	sayHi()

	// Calling function with parameters
	greetPerson("John", "Doe")
	greetPerson("Jane", "Smith")

	// Using function with return value
	fullName := getFullName("Alice", "Johnson")
	fmt.Println("Full name:", fullName)

	// Using function with multiple return values
	name, age := getUserInfo("Bob", "Brown", 25)
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Using function with named return values
	personName, length := getPersonDetails("Charlie", "Wilson")
	fmt.Printf("Person: %s (Length: %d)\n", personName, length)

	// Ignoring return values
	_, justAge := getUserInfo("David", "Lee", 30)
	fmt.Println("Age only:", justAge)
}
