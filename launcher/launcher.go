//+build !darwin,!windows

package launcher

import (
	"fmt"
	"os/exec"
)

// OpenURI opens a given uri in a web browser.
func OpenURI(uri string) {
	cmd := exec.Command("xdg-open", uri)
	err := cmd.Start()
	if err != nil {
		fmt.Errorf("%s\n", err)
	} else {
		err = cmd.Wait()
	}
}
