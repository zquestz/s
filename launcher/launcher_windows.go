//+build windows

package launcher

import (
	"fmt"
	"os/exec"
)

// OpenURI opens a given uri in a web browser.
func OpenURI(binary string, uri string) error {
	if binary != "" {
		return fmt.Errorf("Custom binaries are not supported on Windows.")
	}

	cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", uri)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
