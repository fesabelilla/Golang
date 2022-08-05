package main

import (
	"fmt"
)

type address struct {
	area string
	city string
}

type person struct {
	firstName string
	lastName  string
	location  address
}

func main() {
	person1 := person{
		firstName: "Zahid",
		lastName:  "Hasan",
		location: address{
			area: "baganbari",
			city: "Cumilla",
		},
	}
	fmt.Println(person1)
}
