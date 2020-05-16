package main

import (
	"os"

	"github.com/jawher/mow.cli"
)

// run starts the application.
func run() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	a := cli.App("client", "Fileserver Client")

	a.Command("versions", "show available versions", cmdVersions)
	a.Command("download", "download specific version", cmdDownload)

	return a.Run(os.Args)
}
