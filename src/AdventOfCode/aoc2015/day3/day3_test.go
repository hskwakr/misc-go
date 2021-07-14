package day3

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

func TestMoveSata(t *testing.T) {
	in := []string{
		">",
		"^>v<",
		"^v^v^v^v",
	}
	want := [][]Point{
		{
			Point{0, 0},
			Point{1, 0}, // >
		},
		{
			Point{0, 0},
			Point{0, 1}, // ^
			Point{1, 1}, // >
			Point{1, 0}, // v
			Point{0, 0}, // <
		},
		{
			Point{0, 0},
			Point{0, 1}, // ^
			Point{0, 0}, // v
			Point{0, 1}, // ^
			Point{0, 0}, // v
			Point{0, 1}, // ^
			Point{0, 0}, // v
			Point{0, 1}, // ^
			Point{0, 0}, // v
		},
	}

	for k, v := range in {
		got := MoveSanta(v)
		if !equal(got, want[k]) {
			t.Errorf("\ngot  = %v;\nwant = %v", got, want[k])
		}
	}
}

func TestCountHouses(t *testing.T) {
	in := [][]Point{
		{
			Point{0, 0},
			Point{1, 0}, // >
		},
		{
			Point{0, 0},
			Point{0, 1}, // ^
			Point{1, 1}, // >
			Point{1, 0}, // v
			Point{0, 0}, // <
		},
		{
			Point{0, 0},
			Point{0, 1}, // ^
			Point{0, 0}, // v
			Point{0, 1}, // ^
			Point{0, 0}, // v
			Point{0, 1}, // ^
			Point{0, 0}, // v
			Point{0, 1}, // ^
			Point{0, 0}, // v
		},
	}
	want := []int{
		2,
		4,
		2,
	}

	for k, v := range in {
		got := CountHouses(v)
		if got != want[k] {
			t.Errorf("\ngot  = %v;\nwant = %v", got, want[k])
		}
	}
}
