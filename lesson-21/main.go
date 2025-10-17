package main

import "fmt"

type person struct {
	Name string
	Age  int
}

type circle struct {
	Radius float64
}

func (c circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (p person) GetDetails(string, int) string {
	return fmt.Sprintf("Name %s, Age %s", p.Name, p.Age)
}

type Rectangle struct {
	length float64
	width  float64
}

func (r *Rectangle) AreaRect() float64 { // if the below is pointer reciever , standard way is to make all pointer receiver
	return r.length * r.width
}

func (r *Rectangle) SetLength(l float64) { // if * is removed, no changes
	r.length = l
}

func updateLength(r *Rectangle, l float64) {
	r.SetLength(l)
}

func main() {
	rect := Rectangle{length: 2.0, width: 2.0}
	fmt.Println(rect.AreaRect())
	rect.SetLength(10)
	fmt.Println(rect.AreaRect())

	fmt.Println("\n--------------------------")

	rect2 := Rectangle{length: 3.0, width: 3.0}
	fmt.Println(rect2.AreaRect())
	updateLength(&rect2, 10)
	fmt.Println(rect2.AreaRect()) //value does not update if no & and * , but updates if * and &

}
