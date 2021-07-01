package main

import (
	"fmt"
	"time"

	"github.com/hskwakr/misc-go/sieve-of-eratosthenes/sieve"
)

func main() {
	fmt.Printf("Prime number:%v \n", getPrime(101))
}

func getPrime(size uint) []int {
	s, err := sieve.InitSieve(size)
	if err != nil {
		fmt.Println(err)
	}

	color := s.InitDisplay()

	ch := make(chan *sieve.Sieve)
	go s.Screen(ch)
	for v := range ch {
		time.Sleep(300 * time.Millisecond)
		v.Display(color)
	}

	return s.Prime
}
