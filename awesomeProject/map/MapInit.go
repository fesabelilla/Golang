package main

/*
	Maps is an unordered collection
	The key type must have the operations == and != defined, like string, int, float.
	Hence arrays, slices and structs cannot be used as key type, but pointers and interface types can.
	Structs can be used as a key when we provide Key() or Hash() method

	A map is a reference type and declared
	var map1 map[keytype]valuetype - > var map1 map[int]string
*/

import (
	"fmt"
)

func main() {
	x := map[string]int{
		"Zahid": 26,
		"Hasan": 25,
		"Kamal": 36,
	}
	fmt.Print(x)
	fmt.Println("\n", x["Kamal"])
}
