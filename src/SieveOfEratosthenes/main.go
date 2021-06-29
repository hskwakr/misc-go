package main

import (
	"fmt"
	"math"

	"github.com/fatih/color"
)

func main() {
	var num [101]int
	for k := range num {
		num[k] = k
	}
	var prime []int

	// step 1
	isPrime := initIsPrimeArray(len(num))

	// step 2 3
	sqrt := int(math.Sqrt(float64(len(num))))
	for i := 2; i < sqrt; i++ {
		if isPrime[i] {
			prime = append(prime, i)

			for j := i * i; j < len(num); j += i {
				isPrime[j] = false
			}
		}
	}
	// step 4
	for i := sqrt; i < len(num); i++ {
		if isPrime[i] {
			prime = append(prime, i)
		}
	}

	// display
	c := color.New(color.BgCyan, color.FgWhite)
	for i, v := range num {
		if i == 0 {
			continue
		}
		if i == 1 {
			fmt.Print("   ")
		} else {
			if isPrime[i] {
				c.Printf("%3v", v)
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

	//fmt.Println()
	//fmt.Println("Prime numbers:")
	//for _, v := range prime {
	//	fmt.Printf("%v, ", v)
	//}
}

func initIsPrimeArray(idx int) []bool {
	prime := make([]bool, idx)

	for k := range prime {
		if k <= 1 {
			prime[k] = false
		} else {
			prime[k] = true
		}
	}
	return prime
}
