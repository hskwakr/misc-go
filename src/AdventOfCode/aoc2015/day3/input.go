package day3

import (
	"bufio"
	"io"
	"log"
	"os"
)

// Return input string
func Data(filePath string) string {
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
