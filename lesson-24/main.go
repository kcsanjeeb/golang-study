package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func PrintArea(s Shape) {
	fmt.Printf("Area of %f\n", s.Area())
}

type Square struct {
	Width float64
}

func (s Square) Area() float64 {
	return s.Width * s.Width
}

// embedding interfaces

type Object interface {
	Name() string
	Shape
}

func PrintObj(o Object) {
	fmt.Printf("Area of %f, Name of %s", o.Area(), o.Name())
}

func (r Rectangle) Name() string {
	return "Rectangle"
}

func main() {
	rectangle := Rectangle{Width: 2, Height: 3}
	circle := Circle{Radius: 2}
	square := Square{Width: 3}
	PrintArea(square)

	shapes := []Shape{rectangle, circle}

	// Calculate and print the area of each shape
	for _, shape := range shapes {
		PrintArea(shape)
	}
	fmt.Println("\n---------------------")

	// embedding interfcae
	rectangle2 := Rectangle{Width: 4, Height: 4}
	PrintObj(rectangle2)
}
