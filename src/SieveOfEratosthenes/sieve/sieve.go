package sieve

import (
	"fmt"
	"math"
	"time"

	"github.com/morikuni/aec"
)

// Sieve to find prime numbers.
type Sieve struct {
	size    uint
	num     []int
	isPrime []bool

	Prime []int
}

/*************************************/
/* Error                             */
/*************************************/

// Error for sieve
type SieveError struct {
	When time.Time
	What string
}

func (e *SieveError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

/*************************************/
/* Private method                    */
/*************************************/

// The amount of loops needed to find all prime numbers.
func threshold(size uint) int {
	return int(math.Sqrt(float64(size)))
}

// Table height of the sequence to display.
func row(l uint) int {
	r := math.Ceil(float64(l) / 10)
	return int(r)
}

/*************************************/
/* Public method                     */
/*************************************/

// Initialize a struct Sieve with the size of numbers.
func InitSieve(size uint) (*Sieve, error) {
	// There is no prime numbers when size is under 2
	if size < 2 {
		return nil, &SieveError{
			time.Now(),
			"The size of numbers must be at least 2",
		}
	}

	size++
	s := &Sieve{}
	s.size = size
	s.num = make([]int, int(size))
	s.isPrime = make([]bool, int(size))

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
	return s, nil
}

// Find prime numbers asynchronously.
// When finding a not prime number, the method sends a Sieve struct to the channel.
// This method is intended to call with goroutine.
func (s *Sieve) Screen(ch chan *Sieve, err chan error) {
	// There is no prime numbers when size is under 2
	if s.size < 2 {
		err <- &SieveError{
			time.Now(),
			"The size of numbers must be at least 2",
		}
	}
	// Need to initialize before using this method
	if s.isPrime[0] && s.isPrime[1] {
		err <- &SieveError{
			time.Now(),
			"Need to initialize the structure before using this method",
		}
	}
	close(err)

	for i := 2; i < threshold(s.size); i++ {
		if s.isPrime[i] {
			s.Prime = append(s.Prime, i)

			for j := i * i; j < int(s.size); j += i {
				s.isPrime[j] = false
				ch <- s
			}
		}
	}
	close(ch)

	for i := threshold(s.size); i < int(s.size); i++ {
		if s.isPrime[i] {
			s.Prime = append(s.Prime, i)
		}
	}
}

// Prepare console to display table.
func (s *Sieve) InitDisplay() aec.ANSI {
	for i := 0; i < row(s.size); i++ {
		fmt.Println()
	}

	color := aec.Color3BitB(aec.NewRGB3Bit(255, 85, 0))
	s.Display(color)
	return color
}

// Display a table of number sequence.
func (s *Sieve) Display(color aec.ANSI) {
	//fmt.Print(aec.Up(uint(row(s.size))))
	fmt.Print(aec.Up(10))

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
		if i%10 == 0 || i == int(s.size)-1 {
			fmt.Println()
		}
	}
}
