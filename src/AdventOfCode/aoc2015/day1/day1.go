package day1

import "github.com/hskwakr/misc-go/src/AdventOfCode/aoc2015/day1/input"

// Find right floor
func FindRightFloor() int {
	return MoveSanta(input.Data())
}

func MoveSanta(input string) int {
	result := 0

	for _, v := range input {
		switch v {
		case '(':
			result++
			break
		case ')':
			result--
			break
		}
	}

	return result
}
