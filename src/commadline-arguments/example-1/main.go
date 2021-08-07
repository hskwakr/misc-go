package main

import (
	"flag"
	"fmt"
)

func main() {
	//output := flag.Bool("output", false, "Should there be output?")
	//input := flag.String("input", "file.csv", "The path to the input file")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n\tgo <command> [arguments]\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	//fmt.Println(*output)
	//fmt.Println(*input)

	for n, v := range flag.Args() {
		if n < 2 {
			fmt.Println(n, v)
		}
	}
}
