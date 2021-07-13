package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	readLines2("./input")
}

func readLines1(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		log.Print(strconv.Quote(s.Text()))
	}
	if s.Err() != nil {
		// non-EOF error.
		log.Fatal(s.Err())
	}
}

func readLines2(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	r := bufio.NewReader(f)
	for {
		// line includes '\n'.
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		line = strings.TrimRight(line, "\n")
		log.Print(strconv.Quote(line))
	}
}
