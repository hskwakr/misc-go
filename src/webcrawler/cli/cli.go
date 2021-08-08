package cli

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"

	"github.com/hskwakr/misc-go/src/webcrawler/crawler"
)

const (
	ExitCodeOK               = 0
	ExitCodeParseFlagError   = 1
	ExitCodeArgumentsError   = 1
	ExitCodeApplicationError = 1
)

type CLI struct {
	OutStream, ErrStream io.Writer
}

var (
	url string
)

func (c *CLI) Run(args []string) int {
	if r := c.parse(args); r != 0 {
		return r
	}

	links, err := crawler.GetLinks(url)
	if err != nil {
		log.Println(err)
		return ExitCodeApplicationError
	}
	writeJSON(links)

	return ExitCodeOK
}

func (c *CLI) parse(args []string) int {
	flags := flag.NewFlagSet("webcrawler", flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	// Two arguments are required
	//fmt.Println(len(flags.Args()))
	if len(flags.Args()) < 1 {
		return ExitCodeArgumentsError
	}

	url := flags.Arg(0)
	if !urlValidation(url) {
		return ExitCodeArgumentsError
	}

	return ExitCodeOK
}

func urlValidation(url string) bool {
	r := true

	if len(url) == 0 {
		r = false
	}

	return r
}

func writeJSON(data []crawler.Link) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	_ = ioutil.WriteFile("links.json", file, 0644)
}
