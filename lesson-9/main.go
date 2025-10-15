package main

import "fmt"

func main() {
	// Map declaration examples
	nameAge := map[string]int{
		"Foo": 25,
		"Bar": 20,
	}
	firstNameLastName := map[string]string{
		"Foo": "Bar",
		"Bar": "Foo",
	}
	fmt.Println(nameAge, firstNameLastName)

	// Different ways to create maps
	// var nameAge map[string]int        // Nil map
	// nameAge := map[string]int{}       // Empty map (shorthand)
	// nameAge := make(map[string]int)   // Empty map (make function)
	// var nameAge map[string]int = map[string]int{}  // Empty map (explicit)

	nameAge = map[string]int{}
	nameAge["foo"] = 23
	nameAge["bar"] = 24
	nameAge["foo bar"] = 25
	fmt.Println("Length:", len(nameAge))
	fmt.Println("foo bar:", nameAge["foo bar"])
	fmt.Println("foo:", nameAge["foo"])
	fmt.Println("bar:", nameAge["bar"])

	x := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	// Reading from map
	fmt.Println("Length:", len(x))
	fmt.Println("Map x:", x)
	fmt.Println("Value b:", x["b"])

	// Updating value
	x["b"] = 20
	fmt.Println("Updated b:", x["b"])

	// Comma OK idiom
	nameGrade := map[string]int{
		"john": 10,
		"mike": 20,
		"sawn": 30,
		"buka": 0,
	}
	gradeX, ok := nameGrade["buka"]
	fmt.Println("Grade:", gradeX, "Exists:", ok)
}
