package main

import (
	"fmt"
	"time"

	"github.com/hskwakr/misc-go/src/sieve-of-eratosthenes/sieve"
)

func main() {
	prime := getPrime(99)

	fmt.Println()
	fmt.Printf("Prime number:%v \n", prime)
}

func getPrime(size uint) []int {
	s, err1 := sieve.InitSieve(size)
	if err1 != nil {
		fmt.Println(err1)
		return nil
	}

	color := s.InitDisplay()

	ch := make(chan *sieve.Sieve)
	err2 := make(chan error)
	go s.Screen(ch, err2)

	end := false
	for {
		if end {
			break
		}

		select {
		case e := <-err2:
			if e != nil {
				fmt.Println(e)
				end = true
				break
			}
		case v, ok := <-ch:
			if !ok {
				end = true
				break
			}
			time.Sleep(100 * time.Millisecond)
			v.Display(color)
		}
	}

	return s.Prime
}
