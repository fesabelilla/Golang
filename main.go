// run go : go run FileName.go
package main

import "fmt"

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
	fmt.Printf("userName Type is : %T, userTickets type is : %T", userName, userTickets)
}
