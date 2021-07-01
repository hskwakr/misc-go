package main

import (
	"fmt"
	"time"

	"github.com/hskwakr/misc-go/sieve-of-eratosthenes/sieve"
)

func main() {
	getPrime(101)
}

func getPrime(idx int) []int {
	s, err := sieve.InitSieve(idx)
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
