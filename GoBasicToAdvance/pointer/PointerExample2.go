package main

import (
	"fmt"
)

func main() {
	input := new(int)
	fmt.Println("Before change input value : ", *input)
	changeValue(input)
	fmt.Println("After change input value : ", *input)
}

func changeValue(input *int) {
	*input = 100
}
