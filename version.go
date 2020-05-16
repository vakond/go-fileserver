package main

import (
	"fmt"
	"strings"

	"github.com/jawher/mow.cli"
)

// cmdVersion sets action for command 'version'.
func cmdVersion(c *cli.Cmd) {
	c.Action = func() {
		if err := version(); err != nil {
			throw(err)
		}
	}
}

// version implements command 'version'.
func version() error {
	fmt.Printf("%s %s version %s\n", strings.ToUpper(companyName), appName, appVersion)
	return nil
}
