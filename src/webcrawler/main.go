package main

import (
	"os"

	"github.com/hskwakr/misc-go/src/webcrawler/cli"
)

func main() {
	app := &cli.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(app.Run(os.Args))
}
