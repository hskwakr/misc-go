package day1

import "testing"

func TestMoveSanta(t *testing.T) {
	input := []string{
		"(())",
		"(((",
		"))(",
		")))",
	}
	want := []int{
		0,
		3,
		-1,
		-3,
	}
	for k, v := range input {
		got := MoveSanta(v)
		if got != want[k] {
			t.Errorf("got = %v; want %v", got, want[k])
		}
	}
}
