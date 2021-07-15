package input

import (
	"bufio"
	"io"
	"log"
	"os"
)

// Read text from a file and return it as an array of strings.
// It has a limited length of text per line that can be read.
func ReadLines(filePath string) []string {
	var r []string
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

// Asynchronously read text from a file and return it as an array of strings.
// It has a limited length of text per line that can be read.
func ReadLinesAsync(filePath string, out chan string, err chan error) {
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

// Read a line of text from a text file and return it as a string.
func ReadLine(filePath string) string {
	var r string
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
			return ""
		}
		allLinesProcessed := err == io.EOF && len(line) == 0
		if allLinesProcessed {
			break
		}
		r = string(line)
	}

	return r
}
