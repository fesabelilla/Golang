/*
Reff :  https://www.javatpoint.com/go-interface
Interfaces provide behavior to an object: if something can do this, then it can be used here.
An interface defines a set of abstract methods and does not contain any variable.
*/
package main

import (
	"fmt"
)

type vehicle interface {
	accelerate()
}

func foo(v vehicle) {
	fmt.Println(v)
}

type car struct {
	model string
	color string
}

func (c car) accelerate() {
	fmt.Println("Accelerate........")
}

type toyota struct {
	model string
	color string
	speed int
}

func (t toyota) accelerate() {
	fmt.Println("Toyota, Accelerate fast? ")
}

func main() {
	car1 := car{
		model: "suzuki",
		color: "blue",
	}
	toyota1 := toyota{
		model: "toyota",
		color: "Red",
		speed: 100,
	}
	car1.accelerate()
	toyota1.accelerate()
	foo(car1)
	foo(toyota1)
}
