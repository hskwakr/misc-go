package sieve

import (
	"reflect"
	"testing"
)

/*************************************/
/* Local function                    */
/*************************************/

// Compare two variables
func equal(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

/*************************************/
/* Test                              */
/*************************************/

func TestThreshold(t *testing.T) {
	want := 3
	s, _ := InitSieve(10)
	if want != s.threshold() {
		t.Errorf("s.threshold() = %v; want %v", s.threshold(), want)
	}
}

func TestRow(t *testing.T) {
	want := 10
	s, _ := InitSieve(99)
	if want != s.row() {
		t.Errorf("s.row() = %v; want %v", s.row(), want)
	}
}

func TestInitSieve(t *testing.T) {
	want := []interface{}{
		100,
		true,
		2,
		nil,
	}
	var got *Sieve
	var err error

	got, err = InitSieve(100)
	if err != nil && len(got.num) != want[0] {
		t.Errorf("len(got.num) = %v; want %v", len(got.num), want[0])
	}
	got, err = InitSieve(3)
	if err != nil && got.isPrime[0] && got.isPrime[1] {
		t.Errorf("got.isPrime[0] = %v got.isPrime[1] = %v; want %v", got.isPrime[0], got.isPrime[1], want[1])
	}
	got, err = InitSieve(2)
	if err != nil && got.Prime[0] != want[2] {
		t.Errorf("got.Prime[0] = %v; want %v", got.Prime[0], want[2])
	}
	got, err = InitSieve(1)
	if err == want[3] {
		t.Errorf("err = %v; want %v", err, want[3])
	}
}

func TestScreen_1(t *testing.T) {
	want := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	ch := make(chan *Sieve)
	err := make(chan error)
	got, _ := InitSieve(100)

	go got.Screen(ch, err)
	for {
		e := <-err
		if e != nil {
			t.Errorf("e = %v; want nil", e)
			break
		}

		_, done := <-ch
		if !done {
			break
		}
	}
	if !equal(want, got.Prime) {
		t.Errorf("got.Prime = %v; want = %v", got.Prime, want)
	}
}

func TestScreen_2(t *testing.T) {
	want := [][]bool{
		// false: 4
		{false, false, true, true, false, true, true, true, true, true, true},
		// false: 4, 6
		{false, false, true, true, false, true, false, true, true, true, true},
		// false: 4, 6, 8
		{false, false, true, true, false, true, false, true, false, true, true},
		// false: 4, 6, 8, 10
		{false, false, true, true, false, true, false, true, false, true, false},
		// false: 4, 6, 8, 9, 10
		{false, false, true, true, false, true, false, true, false, false, false},
	}
	ch := make(chan *Sieve)
	err := make(chan error)
	s, _ := InitSieve(10)

	go s.Screen(ch, err)
	end := false
	k := 0
	for {
		if end {
			break
		}

		select {
		case e := <-err:
			if e != nil {
				t.Errorf("e = %v; want nil", e)
				end = true
				break
			}
		case got, ok := <-ch:
			if !ok {
				end = true
				break
			}
			if !equal(want[k], got.isPrime) {
				t.Errorf("got.isPrime = %v; want[%v] = %v", s.isPrime, k, want[k])
				end = true
				break
			}
			k++
		}
	}
}
