package main

import "fmt"

var numOfCars = 2 // Line 1
type Car struct {
	name  string
	model string
	color string
}

var cars = []Car{{name: "Toyota", model: "Corolla", color: "red"}, {name: "Toyota", model: "Innova", color: "gray"}}

func countRedCars() {
	for i := 0; i < len(cars); i++ {
		if cars[i].color == "red" {
			numOfCars += 1                                        // Line 2
			fmt.Println("Inside countRedCars method ", numOfCars) //Line 3
		}
	}
}

func main() {
	fmt.Print(numOfCars)
	countRedCars()
	fmt.Print(numOfCars)
}
