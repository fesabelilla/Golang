package main

import (
	"fmt"
)

// Go Simple Struct - [ used to create user-defined types. ]
type user struct {
	firstName string
	lastName  string
	age       int
}

// Embedded Struct
type eUser struct {
	fName string
	lName string
}
type employee struct {
	eUser
	empId int
}

func (eu eUser) details() {
	fmt.Println(eu, " This is a user")
}

func (emp employee) details() {
	fmt.Println(emp, "This is an Employee")
}

func embeddedStruct() {
	user1 := eUser{
		fName: "Test01",
		lName: "Name01",
	}
	user1.details()

	employee1 := employee{
		eUser: eUser{
			fName: "Test02",
			lName: "Name02",
		},
		empId: 2001,
	}
	employee1.details()
}

func main() {
	x := user{
		firstName: "Zahid",
		lastName:  "Hasan",
		age:       26,
	}
	fmt.Println(x)
	fmt.Println(x.firstName + " " + x.lastName)

	embeddedStruct()
}
