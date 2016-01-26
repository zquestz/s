//+build windows

package launcher

import (
	"fmt"
	"os"
	"os/exec"
)

// OpenURI opens a given uri in a web browser.
func OpenURI(binary string, uri string) {
	selectedBinary := ""

	if binary == "" {
		selectedBinary = "start"
	} else {
		selectedBinary = binary
	}

	command := fmt.Sprintf("%s %s", selectedBinary, uri)
	cmd := exec.Command("cmd", "/k", command)

	// Only attach output to custom binaries.
	if binary != "" {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
