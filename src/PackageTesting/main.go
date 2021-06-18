package main

import (
	"fmt"

	"github.com/hskwakr/misc/package-testing/abs"
	"github.com/hskwakr/misc/package-testing/fibonacci"
)

func main() {
	fmt.Println("abs:")
	fmt.Println(abs.Abs(-1))

	fmt.Println("\nf1:")
	f1()
	fmt.Println("\nf2:")
	f2()
}

func f1() {
	f := fibonacci.Fibonacci1()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func f2() {
	c := make(chan int, 10)
	quit := make(chan int)
	fibonacci.Fibonacci2Async(c, quit)
}
