package sieve

import "testing"

func TestInitSieve(t *testing.T) {
	var got *Sieve
	var err error

	got, err = InitSieve(101)
	if err != nil && len(got.num) != 101 {
		t.Errorf("len(got.num) = %v; want 101", len(got.num))
	}
	got, err = InitSieve(0)
	if err == nil {
		t.Errorf("err = %v; want error", err)
	}
}
