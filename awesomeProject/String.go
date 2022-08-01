package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	var x string = "Hello World"
	fmt.Println(x)
	fmt.Println("Type of variable :", reflect.TypeOf(x), " Length of string : ", len(x))

	// strings.ToLower(variable) || strings.ToUpper(variable)
	// strings.HasPrefix(s,"IN") || strings.HasSuffix(s,"IA")

	// string join
	var arr = []string{"a", "b", "c", "d"}
	// fmt.Println(strings.Join(arr, " | "))
	var newString string = strings.Join(arr, " | ")
	fmt.Println(newString)

	// strings.Repeat(str,4)

	fmt.Println(strings.Contains(newString, "b"))

	//strings.Index(str,"th")
	//strings.Count(str,"e")

	str := "Hi...there"
	fmt.Println(strings.Replace(str, "e", "Z", 2))

	stringSplit()

	//strings.Compare("a", "b")
	//strings.TrimSpace(" \t\n I love my country  \n\t\r\n")
	//strings.ContainsAny("Hello", "A")
	//strings.ContainsAny("Hello", "o & e")

}

func stringSplit() {
	str := "I,love,my,country"
	var arr []string = strings.Split(str, ",")
	fmt.Println(len(arr))
	for i, v := range arr {
		fmt.Println("Index : ", i, " Value : ", v)
	}
}
