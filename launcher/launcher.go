package launcher

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

// OpenURI opens a given uri in a web browser.
func OpenURI(binary string, uri string) error {
	selectedBinary := []string{}

	if binary == "" {
		selectedBinary = append(selectedBinary, DefaultBinary...)
	} else {
		parts, err := splitCommand(binary)
		if err != nil {
			return err
		}

		selectedBinary = append(selectedBinary, parts...)
	}

	if len(selectedBinary) == 0 {
		return fmt.Errorf("no binary specified")
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

// splitCommand splits a command string into arguments.
// Double or single quotes group tokens containing spaces.
// An unquoted command that matches an existing file is
// treated as a single binary path.
func splitCommand(command string) ([]string, error) {
	if !strings.ContainsAny(command, `"'`) {
		if strings.IndexFunc(command, unicode.IsSpace) != -1 {
			if info, err := os.Stat(command); err == nil && !info.IsDir() {
				return []string{command}, nil
			}
		}

		return strings.Fields(command), nil
	}

	args := []string{}
	current := strings.Builder{}
	inToken := false
	var quote rune

	for _, r := range command {
		switch {
		case quote != 0:
			if r == quote {
				quote = 0
			} else {
				current.WriteRune(r)
			}
		case r == '"' || r == '\'':
			quote = r
			inToken = true
		case unicode.IsSpace(r):
			if inToken {
				args = append(args, current.String())
				current.Reset()
				inToken = false
			}
		default:
			current.WriteRune(r)
			inToken = true
		}
	}

	if quote != 0 {
		return nil, fmt.Errorf("unbalanced quotes in %q", command)
	}

	if inToken {
		args = append(args, current.String())
	}

	return args, nil
}
