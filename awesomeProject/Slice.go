/*
slice is a dynamically-sized, segmented view of an underlying array
We define the subset of an array by indicating the start and end index.
*/
package main

import (
	"fmt"
)

func main() {

	fmt.Println("Slice")
	names := [4]string{
		"hasan",
		"kabir",
		"rafiq",
		"fahim",
	}
	fmt.Println("Array : ", names)
	slice1 := names[0:2]
	fmt.Println("Slice1 : ", slice1)
	slice2 := names[1:3]
	fmt.Println("Slice2 : ", slice2)

	slice2[0] = "testText"
	fmt.Println("Slice1 : ", slice1)
	fmt.Println("Slice2 : ", slice2)
	fmt.Println(names)

	//Slice Literal
	sliceLiteral()

	//lower bound and upper bound
	lowerAndUpperBound()

	// Slice Length and Capacity
	lengthAndCapacity()
}

// Slice literal is like an array literal without any length.
func sliceLiteral() {
	literal := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
		{6, true},
	}

	fmt.Println(literal)
}

func lowerAndUpperBound() {
	slice1 := []int{2, 4, 8, 10, 12, 14}
	slice2 := slice1[2:4]
	fmt.Println(slice2)
	slice3 := slice1[:3]
	fmt.Println(slice3)
	slice4 := slice1[2:]
	fmt.Println(slice4)
	fmt.Println(slice1)
}

// length is the number of stored elements and
// the capacity is the number of elements of the underlying array counting from the beginning of the slice.
func lengthAndCapacity() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(slice1)
	// Slice the slice to give it zero length.
	slice2 := slice1[:0]
	printSlice(slice2)
	// Extend its length.
	slice3 := slice1[:4]
	printSlice(slice3)
	// Drop its first two values.
	slice4 := slice1[2:]
	printSlice(slice4)
}
func printSlice(s []int) {
	fmt.Printf("Length = %d and Capacity = %d %v \n", len(s), cap(s), s)
}
