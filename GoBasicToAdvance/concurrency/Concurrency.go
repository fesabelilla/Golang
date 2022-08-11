package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Reff: https://www.javatpoint.com/go-concurrency
	Goroutines : The parts of an application that run concurrently are called goroutines
	Goroutines are lightweight, much lighter than a thread.
*/

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	go func1()
	go func2()
	wg.Wait()
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("Func1 - > ", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("Func2 -> ", i)
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	wg.Done()
}
