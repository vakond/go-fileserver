package main

import (
	"github.com/jawher/mow.cli"
)

// cmdStart sets action for command 'start'.
func cmdStart(c *cli.Cmd) {
	c.Action = func() {
		if err := start(); err != nil {
			throw(err)
		}
	}
}

// start implements command 'start'.
func start() error {
	if err := configure(); err != nil {
		return err
	}

	return serve()
}
