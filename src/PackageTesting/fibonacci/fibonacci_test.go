package fibonacci

import "testing"

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci1()
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := make(chan int, 10)
		q := make(chan int)

		f := func() {
			for i := 0; i < cap(c); i++ {
				<-c
			}
			q <- 0
		}
		go f()
		Fibonacci2(c, q)
	}
}
