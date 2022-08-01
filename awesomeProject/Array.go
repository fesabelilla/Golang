package main

import (
	"fmt"
)

func main() {

	var i, j int
	var arr [3][3]int
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Scan(&arr[i][j])
		}
	}

	fmt.Println()
	var sum int = 0
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			sum += arr[i][j]
			fmt.Print(arr[i][j], " ")
		}
		fmt.Print(" = ", sum)
		sum = 0
		fmt.Println()
	}
}
