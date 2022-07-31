// reff: https://www.javatpoint.com/go-break
package main

import "fmt"

func forLoop(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(n, " * ", i, " = ", n*i)
	}
}

func main() {
	var number int
	fmt.Printf("Enter a number : ")
	fmt.Scanln(&number)
	forLoop(number)

}
