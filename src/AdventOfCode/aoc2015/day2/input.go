package day2

import (
	"bufio"
	"log"
	"os"
)

// Returns string array for input from text file
func Data() []string {
	var r []string
	f, e := os.Open("/home/hskwakr/temporal/misc-go/src/AdventOfCode/aoc2015/day2/input")
	if e != nil {
		log.Fatal(e)
		return []string{}
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		r = append(r, s.Text())
	}
	if s.Err() != nil {
		// non-EOF error.
		log.Fatal(s.Err())
	}

	return r
}

// Returns string array for input from text file
func DataAsync(out chan string, err chan error) {
	f, e := os.Open("/home/hskwakr/temporal/misc-go/src/AdventOfCode/aoc2015/day2/input")
	if e != nil {
		err <- e
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		out <- s.Text()
	}
	if s.Err() != nil {
		// non-EOF error.
		err <- s.Err()
	}
	close(out)
	close(err)
}
