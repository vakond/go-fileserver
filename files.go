package main

import (
	"fmt"
	"io"
	"os"
)

const (
	defaultDirMode  = os.FileMode(0o755)
	defaultFileMode = os.FileMode(0o644)
)

// dirExists checks if a directory exists.
func dirExists(dir string) bool {
	_, err := os.Stat(dir)
	return !os.IsNotExist(err)
}

// fileExists checks if a file exists and is not a directory.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return !os.IsNotExist(err) && !info.IsDir()
}

// closeFile dumps error from Close if any.
func closeFile(file io.Closer) {
	if err := file.Close(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}
}
