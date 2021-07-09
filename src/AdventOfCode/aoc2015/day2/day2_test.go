package day2

import "testing"

func TestConvStrToDemention(t *testing.T) {
	input := []string{
		"2x3x4",
		"1x1x10",
	}
	want := []Dimention{
		{2, 3, 4},
		{1, 1, 10},
	}

	for k, v := range input {
		got, err := ConvStrToDemention(v)
		if err != nil {
			t.Error(err)
		}
		if got != want[k] {
			t.Errorf("got = %v; want %v", got, want[k])
		}
	}
}

func TestCalcPresentSquareFeet(t *testing.T) {
	input := []Dimention{
		{2, 3, 4},
		{1, 1, 10},
	}
	want := []int{
		58,
		43,
	}

	for k, v := range input {
		got := CalcPresentSquareFeet(v)
		if got != want[k] {
			t.Errorf("got = %v; want %v", got, want[k])
		}
	}

}

func TestMin(t *testing.T) {
	input := [][3]int{
		{1, 2, 3},
		{2, 3, 1},
		{3, 1, 2},
	}
	want := []int{
		1,
		1,
		1,
	}

	for k, v := range input {
		got := min(v[0], v[1], v[2])
		if got != want[k] {
			t.Errorf("got = %v; want %v", got, want[k])
		}
	}
}
