package main

import "fmt"

type Engine struct {
	HorsePower int
	Model      string
}

func (e *Engine) Start() {
	fmt.Println("Engine Started")
}

type GPS struct {
	Model string
}

type Car struct {
	Model string
	Engine
	GPS
}

func (c *Car) Drive() {
	fmt.Printf("Driving my %s...\n", c.Model)
}

func main() {
	myCar := Car{
		Model:  "nissan",
		Engine: Engine{HorsePower: 200, Model: "Nissan HR engine"},
		GPS:    GPS{Model: "Garmin"},
	}
	fmt.Println("Car Model :", myCar.Model)
	fmt.Println("Car Engine :", myCar.HorsePower)
	fmt.Println("\n----------------------------")

	myCar.Start()
	myCar.Drive()
	fmt.Println("\n----------------------------")

	fmt.Println("Car Model :", myCar.Model)
	fmt.Println("Car Engine :", myCar.Engine.Model)
	fmt.Println("Car Engine HP :", myCar.Engine.HorsePower)
	fmt.Println("Car GPS Model :", myCar.GPS.Model)

}
