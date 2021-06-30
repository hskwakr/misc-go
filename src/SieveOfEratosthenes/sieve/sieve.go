package sieve

import (
	"fmt"
	"math"

	"github.com/morikuni/aec"
)

type Sieve struct {
	idx     int
	num     []int
	isPrime []bool

	Prime []int
}

/*************************************/
/* Private method                    */
/*************************************/

// The amount of loops needed to find all prime numbers.
func threshold(idx int) int {
	return int(math.Sqrt(float64(idx)))
}

// Table height of the sequence to display.
func row(l int) int {
	return l / 10
}

/*************************************/
/* Public method                     */
/*************************************/

func InitSieve(idx int) *Sieve {
	s := &Sieve{}
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

func (s *Sieve) Screen(ch chan *Sieve) {
	for i := 2; i < threshold(s.idx); i++ {
		if s.isPrime[i] {
			s.Prime = append(s.Prime, i)

			for j := i * i; j < s.idx; j += i {
				s.isPrime[j] = false
				ch <- s
			}
		}
	}
	close(ch)

	for i := threshold(s.idx); i < s.idx; i++ {
		if s.isPrime[i] {
			s.Prime = append(s.Prime, i)
		}
	}
}

func (s *Sieve) InitDisplay() aec.ANSI {
	for i := 0; i < row(s.idx); i++ {
		fmt.Println()
	}

	color := aec.Color3BitB(aec.NewRGB3Bit(255, 85, 0))
	s.Display(color)
	return color
}

func (s *Sieve) Display(color aec.ANSI) {
	fmt.Print(aec.Up(uint(row(s.idx))))

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
