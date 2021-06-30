package main

import (
	"time"

	"github.com/hskwakr/misc-go/sieve-of-eratosthenes/sieve"
)

func main() {
	getPrime(101)
}

func getPrime(idx int) []int {
	s := sieve.InitSieve(idx)
	color := s.InitDisplay()

	ch := make(chan *sieve.Sieve)
	go s.Screen(ch)
	for v := range ch {
		time.Sleep(300 * time.Millisecond)
		v.Display(color)
	}

	return s.Prime
}
