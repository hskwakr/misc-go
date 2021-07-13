package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hskwakr/misc-go/src/AdventOfCode/aoc2015/day1"
	"github.com/hskwakr/misc-go/src/AdventOfCode/aoc2015/day2"
)

func main() {
	filePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("AoC 2015")
	fmt.Printf("Day1: %d\n", day1.FindRightFloor())
	fmt.Printf("Day2: %d\n", day2.CalcTotalSquareFeetAsync(filePath+"/aoc2015/day2/input"))
}
