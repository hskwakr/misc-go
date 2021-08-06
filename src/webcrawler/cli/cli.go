package cli

import (
	"flag"
	"fmt"
	"io"
)

const (
	ExitCodeOK             = 0
	ExitCodeParseFlagError = 1
	ExitCodeArgumentsError = 1
)

type CLI struct {
	OutStream, ErrStream io.Writer
}

func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("webcrawler", flag.ContinueOnError)
	flags.SetOutput(c.ErrStream)

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	// Two arguments are required
	//fmt.Println(len(flags.Args()))
	if len(flags.Args()) < 2 {
		return ExitCodeArgumentsError
	}

	domain := flags.Arg(0)
	url := flags.Arg(1)
	fmt.Println(domain, url)

	return ExitCodeOK
}
