package main

import (
	"fmt"
)

func main() {
	type student struct {
		firstName string
		lastName  string
		age       int
		classes   []string
	}

	// Implicit initialization (commented out)
	// foo := student{
	// 	"code",
	// 	"learn",
	// 	25,
	// 	[]string{"maths", "english"},
	// }

	// Explicit initialization
	foo := student{
		firstName: "John",
		lastName:  "Doe",
	}
	fmt.Println(foo)

	// Zero value struct
	var student1 student
	fmt.Printf("%+v", student1)

	fmt.Printf("\n----------------------------\n")

	// Various initialization methods
	var student2 student
	student2 = student{
		"jane", "Doe", 10, []string{"maths", "english"},
	}
	fmt.Printf("%+v", student2)

	fmt.Printf("\n----------------------------\n")

	student3 := student{
		"jack", "Doe", 9, []string{"maths", "english"},
	}
	fmt.Printf("%+v", student3)

	fmt.Printf("\n----------------------------\n")

	student4 := student{
		firstName: "Jade",
		lastName:  "Doe",
		age:       9,
		classes:   []string{"maths", "english"},
	}
	fmt.Printf("%+v", student4)
	fmt.Println("First name of student 4:", student4.firstName)

	fmt.Printf("\n----------------------------\n")

	// Modifying struct fields
	student4.classes = append(student4.classes, "arts")
	fmt.Printf("%+v", student4)

	fmt.Printf("\n----------------------------\n")

	// Anonymous struct
	guardian := struct {
		firstName string
		lastName  string
	}{
		firstName: "Alex",
		lastName:  "Moon",
	}
	fmt.Printf("%+v", guardian)
}
