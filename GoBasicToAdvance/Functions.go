package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
}

func (employee Employee) fullName() {
	fmt.Println(employee.firstName + " " + employee.lastName)
}

func sumOfTwoNumber(a int, b int) int {
	return a + b
}

func addAll(args ...int) (int, int) {
	finalAddValue := 0
	finalSubValue := 0

	for _, value := range args {
		finalAddValue += value
		finalSubValue -= value
	}
	return finalAddValue, finalSubValue
}

func main() {
	emp := Employee{
		firstName: "Zahid",
		lastName:  " Hasan",
	}
	emp.fullName()

	var sum int
	sum = sumOfTwoNumber(5, 6)
	fmt.Println(sum)

	fmt.Println("Go function multiple return")
	fmt.Println(addAll(10, 15, 20, 25, 30))
}
