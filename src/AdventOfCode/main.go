package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hskwakr/misc-go/src/AdventOfCode/aoc2015/day1"
	"github.com/hskwakr/misc-go/src/AdventOfCode/aoc2015/day2"
	"github.com/hskwakr/misc-go/src/AdventOfCode/aoc2015/day3"
)

func main() {
	// Get working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("AoC 2015")
	fmt.Printf("Day1: %d\n", day1.FindRightFloor(wd+"/aoc2015/day1/input"))
	fmt.Printf("Day2: %d\n", day2.CalcTotalSquareFeetAsync(wd+"/aoc2015/day2/input"))
	fmt.Printf("Day3: %d\n", day3.Count(wd+"/aoc2015/day3/input"))
}
