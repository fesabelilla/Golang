package main

// Panic -> Defer -> Recover

import "fmt"

func x() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover in x ", r)
		}
	}()
	fmt.Println("Executing X......")
	fmt.Println("Calling Y.")
	y(0)
	fmt.Println("Returned normally form Y.")
}

func y(i int) {
	fmt.Println("Executing Y......")
	if i > 2 {
		fmt.Println("Panicking !!!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in Y", i)
	fmt.Println("Printing in y", i)
	y(i + 1)
}

func main() {
	fmt.Println("Calling X from main.")
	x()
	fmt.Println("Returned from x.")
}
