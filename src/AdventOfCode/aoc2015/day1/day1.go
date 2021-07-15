package day1

import "github.com/hskwakr/misc-go/src/AdventOfCode/util/input"

// Find right floor
func FindRightFloor(filePath string) int {
	return MoveSanta(input.ReadLine(filePath))
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
