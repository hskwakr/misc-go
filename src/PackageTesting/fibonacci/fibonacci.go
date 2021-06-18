package fibonacci

import "fmt"

func Fibonacci1() func() int {
	x := 0
	y := 1
	result := 0

	return func() int {
		result = x
		x, y = y, x+y

		return result
	}
}

func Fibonacci2Async(c, quit chan int) {
	f := func() {
		for i := 0; i < cap(c); i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}

	go f()
	Fibonacci2(c, quit)
}

func Fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			return
		}
	}
}
