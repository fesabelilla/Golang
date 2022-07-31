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

func main() {
	emp := Employee{
		firstName: "Zahid",
		lastName:  " Hasan",
	}
	emp.fullName()
}
