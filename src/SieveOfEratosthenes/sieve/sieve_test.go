package sieve

import "testing"

/*************************************/
/* Local function                    */
/*************************************/

// Compeare two int slice.
func equal(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for k, v1 := range s1 {
		if s2[k] != v1 {
			return false
		}
	}
	return true
}

/*************************************/
/* Test                              */
/*************************************/

func TestInitSieve(t *testing.T) {
	var got *Sieve
	var err error

	got, err = InitSieve(101)
	if err != nil && len(got.num) != 101 {
		t.Errorf("len(got.num) = %v; want 101", len(got.num))
	}
	got, err = InitSieve(2)
	if err != nil && len(got.num) != 2 {
		t.Errorf("len(got.num) = %v; want 2", len(got.num))
	}
	got, err = InitSieve(3)
	if err != nil && got.isPrime[0] && got.isPrime[1] {
		t.Errorf("got.isPrime[0] = %v got.isPrime[1] = %v; want both true", got.isPrime[0], got.isPrime[1])
	}
	got, err = InitSieve(0)
	if err == nil {
		t.Errorf("err = %v; want error", err)
	}
}

func TestScreen(t *testing.T) {
	want := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	ch := make(chan *Sieve)
	err := make(chan error)
	got, _ := InitSieve(101)

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
