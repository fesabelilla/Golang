package main

/*
	Go panic is a mechanism by which we handle error situations.
	Panic can be used to abort a function execution.
	When a function calls panic, its execution stops and the control flows to the associated deferred function.
	The caller of this function also gets terminated and caller's deferred function gets executed (if present any).
	This process continues till the program terminates.
	Now the error condition is reported.
	This termination sequences is called panicking and can be controlled by the built-in function recover.
*/

import "os"

func main() {
	panic("Error Situation")
	_, err := os.Open("/temp/file")
	if err != nil {
		panic(err)
	}
}
