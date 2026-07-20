package launcher

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

// TestSplitCommand checks command splitting and quoting rules.
func TestSplitCommand(t *testing.T) {
	cases := []struct {
		command  string
		expected []string
		valid    bool
	}{
		{"w3m", []string{"w3m"}, true},
		{"chromium --incognito", []string{"chromium", "--incognito"}, true},
		{"chromium  --incognito", []string{"chromium", "--incognito"}, true},
		{"", []string{}, true},
		{"   ", []string{}, true},
		{`"C:\Program Files\Mozilla Firefox\firefox.exe" -private-window`, []string{`C:\Program Files\Mozilla Firefox\firefox.exe`, "-private-window"}, true},
		{`'/Applications/Google Chrome' --foo`, []string{"/Applications/Google Chrome", "--foo"}, true},
		{`browser --profile="my profile"`, []string{"browser", "--profile=my profile"}, true},
		{`"unbalanced`, nil, false},
		{`browser 'unbalanced`, nil, false},
	}

	for _, c := range cases {
		args, err := splitCommand(c.command)
		if c.valid {
			if err != nil {
				t.Errorf("splitCommand(%q) returned error: %s", c.command, err)
				continue
			}
			if !slices.Equal(args, c.expected) {
				t.Errorf("splitCommand(%q) = %v, want %v", c.command, args, c.expected)
			}
		} else {
			if err == nil {
				t.Errorf("splitCommand(%q) should return an error.", c.command)
			}
		}
	}
}

// TestSplitCommandPath checks that an unquoted path with
// spaces matching an existing file is a single binary.
func TestSplitCommandPath(t *testing.T) {
	path := filepath.Join(t.TempDir(), "my browser")
	if err := os.WriteFile(path, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatalf("failed to write test binary: %s", err)
	}

	args, err := splitCommand(path)
	if err != nil {
		t.Fatalf("splitCommand(%q) returned error: %s", path, err)
	}

	if len(args) != 1 || args[0] != path {
		t.Errorf("splitCommand(%q) = %v, want single path", path, args)
	}
}
