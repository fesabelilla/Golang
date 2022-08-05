/*
A closure is a function which refers reference variable
from outside its body. The function may access and assign
to the referenced variables.
*/

package main

import (
	"fmt"
)

func main() {

	var number int = 10
	sequareNumber := func() int {
		number *= number
		return number
	}

	fmt.Println(sequareNumber())
	fmt.Println(sequareNumber())
}
