package main

import "fmt"

type Printer struct {
}

func (p *Printer) Print(m string) {
	fmt.Println(m)
}

func main() {
	printer := &Printer{}

	printFunction := printer.Print

	printFunction("Hello world")
}
