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

func main() {
	//var number int
	//fmt.Printf("Enter a number : ")
	//fmt.Scanln(&number)
	//forLoop(number)
	//
	//forCondition()

	rangeConstruct()
}
