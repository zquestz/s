//+build darwin

package launcher

import (
	"os"
	"os/exec"
)

// OpenURI opens a given uri in a web browser.
func OpenURI(binary string, uri string) error {
	selectedBinary := ""

	if binary == "" {
		selectedBinary = "open"
	} else {
		selectedBinary = binary
	}

	cmd := exec.Command(selectedBinary, uri)

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
