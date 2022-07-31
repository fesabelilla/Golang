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
	return a + b;
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
}
