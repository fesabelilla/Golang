// reff: https://www.javatpoint.com/go-break
package main

import "fmt"

func forLoop(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(n, " * ", i, " = ", n*i)
	}
}

func forCondition() {
	sum := 1
	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}
}

func rangeConstruct() {
	nums := []int{1, 2, 3, 4, 5, 6}
	sum := 0
	for _, value := range nums {
		sum += value
	}
	fmt.Println("Sum : ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index : ", i)
		}
	}

	kvs := map[string]string{"1": "mango", "2": "apple", "3": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	for k := range kvs {
		fmt.Println("Key : ", k)
	}

	for i, c := range "hello" {
		fmt.Println(i, c)
	}
}

func gotoStatement() {
	var input int

Loop:
	fmt.Println("Not eligible to vote !!")
	fmt.Print("Enter your age : ")
	fmt.Scanln(&input)
	if input <= 17 {
		goto Loop
	} else {
		fmt.Print("you can vote")
	}
}

func breakStatement() {
	var a int = 1

	for a < 10 {
		fmt.Println("Value of a is : ", a)
		a++
		if a > 5 {
			break
		}
	}
}

func continueStatement() {
	var a int = 1
	for a < 10 {
		if a == 5 {
			/* skip the iteration */
			a = a + 1
			continue
		}
		fmt.Printf("value of a: %d\n", a)
		a++
	}
}

func main() {
	var number int
	fmt.Printf("Enter a number : ")
	fmt.Scanln(&number)
	forLoop(number)

	forCondition()

	rangeConstruct()

	gotoStatement()

	breakStatement()

	continueStatement()
}
