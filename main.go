package main

import (
	"log"
)

func main() {
	if err := run(); err != nil {
		format := "%v"
		if config.debug {
			format = "%+v"
		}
		log.Fatalf("Error: "+format, err)
	}
}
