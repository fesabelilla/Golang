package main

import (
	"fmt"
)

func checkEvenOdd(number int) {
	fmt.Print(number)
	if number%2 == 0 {
		fmt.Println(" is even")
	} else {
		fmt.Printf(" is odd")
	}
}

func main() {
	fmt.Printf("Enter a number : ")
	var number int
	fmt.Scanln(&number)
	checkEvenOdd(number)
}
