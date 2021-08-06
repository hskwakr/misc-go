package main

import (
	"flag"
	"fmt"
)

func init() {
	flag.Parse()
	fmt.Println(flag.Args())
}

func main() {
}
