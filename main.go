//go mod init projectName
// run go : go run FileName.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	fmt.Println(conferenceName)

	var userName string
	var userTickets int
	//pointer - memory location ( &VariableName)
	fmt.Println("Enter your user name :")
	fmt.Scan(&userName)
	fmt.Println("Enter your ticket count : ")
	fmt.Scan(&userTickets)
	// variable value
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

	// variable type
	fmt.Printf("userName Type is : %T, userTickets type is : %T\n", userName, userTickets)

	//ArraysAndSlices
	//array
	var names [50]string
	names[0] = "zahid"
	names[1] = "hasan"

	fmt.Printf("The whole array: %v\n", names)
	fmt.Printf("The first value : %v\n", names[0])
	fmt.Printf("Array Type : %T\n", names)
	fmt.Printf("Array Length : %v\n", len(names))

	//slice
	var slice []string
	slice = append(slice, userName)
	fmt.Printf("The whole slice : %v\n", slice)
	fmt.Printf("The first value : %v\n", slice[0])
	fmt.Printf("Slice type : %T\n", slice)
	fmt.Printf("Size of the slice : %v\n", len(slice))

	//loop
	firstNames := []string{} //declare a slice
	firstNames = append(firstNames, "zahid hasan")
	firstNames = append(firstNames, "Sayeed hasan")
	outPut := []string{}
	for _, firstName := range firstNames {
		var names = strings.Fields(firstName)
		outPut = append(outPut, names[0])
	}

	fmt.Printf("Only first name :  %v\n", outPut)
}
