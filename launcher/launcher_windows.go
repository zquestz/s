//+build windows

package launcher

import (
	"fmt"
	"os"
	"os/exec"
)

// OpenURI opens a given uri in a web browser.
func OpenURI(binary string, uri string) error {
	if binary != "" {
		return fmt.Errorf("Custom binaries are not supported on Windows.")
	}

	cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", uri)

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
