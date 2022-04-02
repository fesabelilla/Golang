package main

import "fmt"

func main() {

	//Multiple Variable Declaration
	/*
		var a, b = 1, "test01"
		c, d := 2, "test02"

		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(d)
	*/

	//Go Variable Declaration in a Block
	
	var (
		a int
		b int    = 1
		c string = "test"
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
