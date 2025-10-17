package main

import (
	"errors"
	"fmt"
)

type student struct {
	firstname string
	lastname  string
}

func main() {
	var intPtr *int     // asterik in front of datatype denotes that intPtr variable is the pointer to that datatype
	fmt.Println(intPtr) // <nil>
	// fmt.Println(*intPtr) //error

	age := 10
	intPtr = &age
	fmt.Println(intPtr) // 0x1400000e0f8

	// Dereferencing
	fmt.Println(*intPtr) // 10 // but when you use * with pointer it gives you value stored at that address

	name := "Learn Golang"
	namePtr := &name
	fmt.Println(namePtr)  // 0x1400008e040
	fmt.Println(*namePtr) // Learn Golang
	fmt.Println("\n------------------------")

	// MUTABLE
	var a int = 10
	a = 20
	fmt.Println(a)

	// IMMUTABLE
	const b = 30
	// b = 25
	fmt.Println(b)

	fmt.Println("\n------------------------")

	c := 10
	increment(c)
	fmt.Println(c) // value is still 10

	d := 10
	incrementActual(&d) // passing the value
	fmt.Println(d)      // value is still 10

	fmt.Println("\n------------------------")

	s := student{
		firstname: "John",
		lastname:  "Python",
	}
	fmt.Println(s)
	updateLastname(&s, "Golang")
	fmt.Println(s)

	fmt.Println("\n------------------------")
	k := student{
		firstname: "John",
		lastname:  "Python",
	}

	previouslastname, err := updateLastnamewitherr(&k, "java")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(previouslastname)

}

func updateLastnamewitherr(s *student, newlastname string) (*string, error) {
	if newlastname == "" {
		return nil, errors.New("empty new last name")
	}
	previous := s.lastname
	s.lastname = newlastname
	return &previous, nil
}

func updateLastname(s *student, newlastname string) {
	// in this case you dont have to use *s like that of incrementActual function, why explain
	s.lastname = newlastname
}

func increment(x int) {
	x++ // perfoming operation in the copy and not on the actual variable that is on the memory
}

func incrementActual(x *int) {
	*x++ //dereferencing the memory address
}
