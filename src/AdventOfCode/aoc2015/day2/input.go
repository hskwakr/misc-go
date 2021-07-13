package day2

import (
	"bufio"
	"log"
	"os"
)

// Returns string array for input from text file
func Data(filePath string) []string {
	var r []string
	//filePath, e1 := os.Getwd()
	//if e1 != nil {
	//	log.Fatal(e1)
	//	return []string{}
	//}

	//f, e2 := os.Open(filePath + "/aoc2015/day2/input")
	f, e := os.Open(filePath)
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
func DataAsync(filePath string, out chan string, err chan error) {
	f, e := os.Open(filePath)
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
