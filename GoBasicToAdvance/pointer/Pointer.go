package main

/*
	pointer is a variable that stores the address of another variable
	var var_name *var-type
	newly declared pointer which has not been assigned to a variable has the nil value.
	With pointers, we can pass a reference to a variable (for example, as a parameter to a function),
	instead of passing a copy of the variable which can reduce memory usage and increase efficiency.
*/

import (
	"fmt"
)

func changeVariableValue(x *int) {
	*x = 0
}

func main() {
	x := 10
	changeVariableValue(&x)
	fmt.Println(x)
}
