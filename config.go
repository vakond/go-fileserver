package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	companyName = "vakond"
	appName     = "fileserver"
	appVersion  = "0.0.1"

	configHeader = `# This is the fileserver configuration file.
# Format:
#     port      <whitespace> <number>
#     <version> <whitespace> <filename>
# Example:
#     port      8090
#     v1.0.0    /home/user/.config/vakond/fileserver/0.zip
#     v1.1.0    /home/user/.config/vakond/fileserver/1.zip
#     v1.2.0    /home/user/.config/vakond/fileserver/2.zip
`
)

// Config stores the configuration data.
type Config struct {
	verbose  bool
	debug    bool
	port     int
	versions map[string]string
}

var (
	config       Config
	versionRegex *regexp.Regexp
)

func init() {
	versionRegex = regexp.MustCompile(`^v\d+\.\d+\.\d+$`)
}

// correctVersion checks format of version string.
func correctVersion(ver string) bool {
	return versionRegex.MatchString(ver)
}

// configure reads config from the configuration file.
func configure() error {
	configFilename, err := getConfigFilename()
	if err != nil {
		return err
	}

	if !fileExists(configFilename) {
		if err = createConfig(configFilename); err != nil {
			return err
		}
	}

	return readConfig(configFilename)
}

// readConfig fills the config structure data from a file.
func readConfig(filename string) error {
	config.versions = make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer closeFile(file)

	scanner := bufio.NewScanner(file)
	const comment = '#'
	const port = "port"
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || line[0] == comment {
			continue
		}
		pair := strings.Fields(line)
		if len(pair) != 2 {
			return fmt.Errorf("invalid config format: '%s'", line)
		}
		if pair[0] == port {
			value, err := strconv.Atoi(pair[1])
			if err != nil {
				return fmt.Errorf("invalid port format: '%s'", line)
			}
			config.port = value
			continue
		}
		if !correctVersion(pair[0]) {
			return fmt.Errorf("invalid version format: '%s'", line)
		}
		if !fileExists(pair[1]) {
			return fmt.Errorf("file missing: '%s'", line)
		}
		_, found := config.versions[pair[0]]
		if found {
			return fmt.Errorf("duplicated version: '%s'", line)
		}
		config.versions[pair[0]] = pair[1]
	}

	return scanner.Err()
}

// createConfig generates new empty config file.
func createConfig(filename string) error {
	filename = filepath.Base(filename)

	dir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	if !dirExists(dir) {
		if err = os.Mkdir(dir, defaultDirMode); err != nil {
			return err
		}
	}

	dir = filepath.Join(dir, companyName)
	if !dirExists(dir) {
		if err = os.Mkdir(dir, defaultDirMode); err != nil {
			return err
		}
	}

	dir = filepath.Join(dir, appName)
	if !dirExists(dir) {
		if err = os.Mkdir(dir, defaultDirMode); err != nil {
			return err
		}
	}

	file := filepath.Join(dir, filename)
	if !fileExists(file) {
		newConfig, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, defaultFileMode)
		if err != nil {
			return err
		}
		defer closeFile(newConfig)
		_, err = newConfig.WriteString(configHeader)
		if err != nil {
			return err
		}
	}

	log.Println("Created", file)

	return nil
}

// getConfigFilename returns standard absolute name of the config file.
func getConfigFilename() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, companyName, appName, "config"), nil
}
