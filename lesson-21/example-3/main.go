package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) GetDetails() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

type Rectangle struct {
	length int
	width  int
}

func (r *Rectangle) SetLength(newlength int) {
	r.length = newlength
}

func main() {
	f1 := Person.GetDetails

	// create an

	p := Person{Name: "Alice", Age: 12}

	fmt.Println(f1(p)) // Name: Alice, Age: 12

	// method expression for the setlength method of rectangle

	f2 := (*Rectangle).SetLength

	r := &Rectangle{length: 5, width: 3}

	f2(r, 1)

	fmt.Println(r.length)

}
