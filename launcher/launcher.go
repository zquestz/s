package launcher

import (
	"os"
	"os/exec"
	"strings"
)

// OpenURI opens a given uri in a web browser.
func OpenURI(binary string, uri string) error {
	selectedBinary := []string{}

	if binary == "" {
		selectedBinary = append(selectedBinary, DefaultBinary...)
	} else {
		selectedBinary = append(selectedBinary, strings.Split(binary, " ")...)
	}

	selectedBinary = append(selectedBinary, uri)

	cmd := exec.Command(selectedBinary[0], selectedBinary[1:]...)

	// Only attach output to custom binaries.
	if binary != "" {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
