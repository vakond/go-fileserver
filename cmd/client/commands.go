package main

import (
	"github.com/jawher/mow.cli"
)

// cmdVersions is handler of command 'versions'.
func cmdVersions(c *cli.Cmd) {
	c.Action = func() {
		if err := versions(); err != nil {
			throw(err)
		}
	}
}

// cmdDownload is handler of command 'download'.
func cmdDownload(c *cli.Cmd) {
	ver := c.StringArg("VERSION", "", "version")
	c.Spec = "VERSION"
	c.Action = func() {
		if err := download(*ver); err != nil {
			throw(err)
		}
	}
}
