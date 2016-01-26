package launcher

import (
	"os"
	"os/exec"
	"runtime"
)

// OpenURI opens a given uri in a web browser.
func OpenURI(binary string, uri string) error {
	var cmd *exec.Cmd

	if binary == "" {
		switch runtime.GOOS {
		case "darwin":
			cmd = exec.Command("open", uri)
		case "windows":
			cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", uri)
		default:
			cmd = exec.Command("xdg-open", uri)
		}
	} else {
		cmd = exec.Command(binary, uri)
		// Only attach output to custom binaries.
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}
