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
	fmt.Println("Init Map : ")
	x := map[string]int{
		"Zahid": 26,
		"Hasan": 25,
		"Kamal": 36,
	}
	fmt.Print(x)
	fmt.Println("\n The value is : ", x["Kamal"])

	fmt.Println("Map insert and Update : ")
	insertAndUpdate(x)

	fmt.Println("key is present in the table with two value example")
	keyIsPresentWithTwoValue(x)
}

/*
update and insert :

	If the map does not contain the provided key the insert operation will takes place
	If the key is present in the map then update operation takes place.
*/
func insertAndUpdate(x map[string]int) {
	fmt.Println("Insert value : [Halim] = 32 ")
	x["Halim"] = 32
	fmt.Println(x)

	fmt.Println("Update value : [Zahid] = 25 ")
	x["Zahid"] = 25
	fmt.Println(x)

	fmt.Println("Map delete operation by Key : [Kamal] ")
	delete(x, "Kamal")
	fmt.Println(x)
}

// key is present in the table with two value example
func keyIsPresentWithTwoValue(x map[string]int) {
	v, ok := x["Zahid"]
	fmt.Println("The value is : ", v, " is Present : ", ok)

	i, j := x["Test"]
	fmt.Println("The value is : ", i, " is Present : ", j)

}