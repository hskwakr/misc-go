package main

import (
	"fmt"
	"math"
	"time"

	"github.com/morikuni/aec"
)

func main() {
	getPrime(101)
}

type sieve struct {
	idx     int
	num     []int
	prime   []int
	isPrime []bool
}

func getPrime(idx int) []int {
	s := initSieve(idx)
	color := s.initDisplay()

	ch := make(chan *sieve)
	go s.screen(ch)
	for v := range ch {
		time.Sleep(300 * time.Millisecond)
		v.display(color)
	}

	return s.prime
}

func initSieve(idx int) *sieve {
	s := &sieve{}
	s.idx = idx
	s.num = make([]int, idx)
	s.isPrime = make([]bool, idx)

	for k := range s.num {
		s.num[k] = k
	}

	for k := range s.isPrime {
		if k <= 1 {
			s.isPrime[k] = false
		} else {
			s.isPrime[k] = true
		}
	}
	return s
}

func threshold(idx int) int {
	return int(math.Sqrt(float64(idx)))
}

func (s *sieve) screen(ch chan *sieve) {
	for i := 2; i < threshold(s.idx); i++ {
		if s.isPrime[i] {
			s.prime = append(s.prime, i)

			for j := i * i; j < s.idx; j += i {
				s.isPrime[j] = false
				ch <- s
			}
		}
	}
	close(ch)

	for i := threshold(s.idx); i < s.idx; i++ {
		if s.isPrime[i] {
			s.prime = append(s.prime, i)
		}
	}
}

func maxLen(l int) int {
	return l / 10
}

func (s *sieve) initDisplay() aec.ANSI {
	for i := 0; i < maxLen(s.idx); i++ {
		fmt.Println()
	}

	color := aec.Color3BitB(aec.NewRGB3Bit(255, 85, 0))
	s.display(color)
	return color
}

func (s *sieve) display(color aec.ANSI) {
	fmt.Print(aec.Up(uint(maxLen(s.idx))))

	for i, v := range s.num {
		if i == 0 {
			continue
		}
		if i == 1 {
			fmt.Print("   ")
		} else {
			if s.isPrime[i] {
				fmt.Printf(color.Apply("%3v"), v)
			} else {
				fmt.Printf("%3v", v)
			}
		}

		if i%10 > 0 {
			fmt.Print("|")
		}
		if i%10 == 0 {
			fmt.Println()
		}
	}
}
