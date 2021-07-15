package day2

import (
	"fmt"
	"time"

	"github.com/hskwakr/misc-go/src/AdventOfCode/util/input"
)

// A box which is a perfect right rectangular prism
type Dimention struct {
	l int
	w int
	h int
}

// Error for sieve
type DimentionError struct {
	When time.Time
	What string
}

func (e *DimentionError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// Calculate total square feet should elves order
func CalcTotalSquareFeet(filePath string) int {
	total := 0
	in := input.ReadLines(filePath)
	for _, v := range in {
		total += CalcArea(v)
	}

	return total
}

// Asynchronously calculate the total area that the elves should order
func CalcTotalSquareFeetAsync(filePath string) int {
	total := 0
	in := make(chan string)
	err := make(chan error)

	go input.ReadLinesAsync(filePath, in, err)
	end := false
	for {
		if end {
			break
		}

		select {
		case e := <-err:
			if e != nil {
				fmt.Println(e)
				end = true
				break
			}
		case v, ok := <-in:
			if !ok {
				end = true
				break
			}

			total += CalcArea(v)
		}
	}
	return total
}

// Calculate an area of paper for a present square feet
func CalcArea(in string) int {
	d, err := ConvStrToDimention(in)
	if err != nil {
		fmt.Println(err)
	}
	return CalcPresentSquareFeet(d)
}

// Convert string input value into Dimention
func ConvStrToDimention(input string) (Dimention, error) {
	var result Dimention

	format := "%dx%dx%d"
	_, err := fmt.Sscanf(input, format, &result.l, &result.w, &result.h)
	if err != nil {
		return Dimention{}, &DimentionError{
			time.Now(),
			err.Error(),
		}
	}

	// Input values should be natural numbers
	if result.l < 0 && result.w < 0 && result.h < 0 {
		return Dimention{}, &DimentionError{
			time.Now(),
			"input values should be natural numbers",
		}
	}

	return result, nil
}

// Calculate a square feet of a present
func CalcPresentSquareFeet(d Dimention) int {
	result := 0

	// 2*l*w + 2*w*h + 2*h*l
	lw := d.l * d.w
	wh := d.w * d.h
	hl := d.h * d.l
	result = 2*lw + 2*wh + 2*hl
	result += min(lw, wh, hl)

	return result
}

// Return minimal value
func min(a, b, c int) int {
	result := a
	if result > b {
		result = b
	}
	if result > c {
		result = c
	}
	return result
}
