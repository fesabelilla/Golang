package main

//  rune is an alias for type int32
// 	rune literals are an integer value.
// 	it will provide the ASCII value of the character.

import (
	"fmt"
	"reflect"
)

func main() {
	rune := 'A'
	fmt.Printf("%d\n", rune)
	fmt.Println(reflect.TypeOf(rune))
}
