package main

import (
	"os"
	"strings"

	"github.com/jawher/mow.cli"
)

// run dispatches commands.
func run() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	a := cli.App(appName, strings.ToUpper(companyName)+" "+appName)

	a.BoolOptPtr(&config.verbose, "v verbose", false, "show additional information")
	a.BoolOptPtr(&config.debug, "d debug", false, "show debug information")

	a.Command("start", "start operation", cmdStart)
	a.Command("version", "show version", cmdVersion)

	return a.Run(os.Args)
}

// throw calls panic when it means throwing exception (i.e. recovered later).
func throw(err error) {
	panic(err) // will be handled by mow.cli
}
