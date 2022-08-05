package main

import (
	"errors"
	"fmt"
	"math"
)

/*
	for error handling go use defer-panic-and-recover mechanism.
	Return an error object. If no error occurred the value will nill
*/

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}

func main() {
	result, err := Sqrt(-64)
	printErrorOrNot(result, err)

	result, err = Sqrt(64)
	printErrorOrNot(result, err)
}

func printErrorOrNot(result float64, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
