package main

// examine its own structure
//Reflect can be used to investigate types and variables at runtime,

import (
	"fmt"
	"reflect"
)

func main() {

	age := 27.5
	fmt.Printf("%T\n", age)
	fmt.Println(reflect.TypeOf(age))
}
