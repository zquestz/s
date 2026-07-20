package providers

import "testing"

// setupTestProviders installs a fixed set of providers
// and restores the original providers after the test.
func setupTestProviders(t *testing.T) {
	t.Helper()

	orig := Providers
	Providers = map[string]Provider{}

	t.Cleanup(func() {
		Providers = orig
	})

	AddProvider("alpha", &CustomProvider{Name: "alpha", URL: validURL, TagList: []string{"linux", "code"}})
	AddProvider("beta", &CustomProvider{Name: "beta", URL: validURL, TagList: []string{"code"}})
	AddProvider("gamma", &CustomProvider{Name: "gamma", URL: validURL})
}

// TestRegion checks region codes parsed from locales.
func TestRegion(t *testing.T) {
	cases := []struct {
		locale   string
		expected string
	}{
		{"en_US", "US"},
		{"de_DE", "DE"},
		{"en-GB", "GB"},
		{"fr_FR", "FR"},
		{"", "US"},
		{"not a locale", "US"},
	}

	for _, c := range cases {
		if r := Region(c.locale); r != c.expected {
			t.Errorf("Region(%q) = %q, want %q", c.locale, r, c.expected)
		}
	}
}

// TestLanguage checks language codes parsed from locales.
func TestLanguage(t *testing.T) {
	cases := []struct {
		locale   string
		expected string
	}{
		{"en_US", "en"},
		{"de_DE", "de"},
		{"es", "es"},
		{"", "en"},
		{"not a locale", "en"},
	}

	for _, c := range cases {
		if l := Language(c.locale); l != c.expected {
			t.Errorf("Language(%q) = %q, want %q", c.locale, l, c.expected)
		}
	}
}

// TestExpandProvider checks exact and prefix matching of provider names.
func TestExpandProvider(t *testing.T) {
	setupTestProviders(t)

	cases := []struct {
		input    string
		expected string
		valid    bool
	}{
		{"alpha", "alpha", true},
		{"al", "alpha", true},
		{"b", "beta", true},
		{"g", "gamma", true},
		{"zeta", "", false},
		{"c++", "", false},
		{"", "", false},
	}

	for _, c := range cases {
		name, err := ExpandProvider(c.input)
		if c.valid {
			if err != nil {
				t.Errorf("ExpandProvider(%q) returned error: %s", c.input, err)
			}
			if name != c.expected {
				t.Errorf("ExpandProvider(%q) = %q, want %q", c.input, name, c.expected)
			}
		} else {
			if err == nil {
				t.Errorf("ExpandProvider(%q) should return an error.", c.input)
			}
		}
	}
}

// TestExpandTag checks exact and prefix matching of tag names.
func TestExpandTag(t *testing.T) {
	setupTestProviders(t)

	cases := []struct {
		input    string
		expected string
		valid    bool
	}{
		{"code", "code", true},
		{"c", "code", true},
		{"li", "linux", true},
		{"video", "", false},
		{"c++", "", false},
		{"", "", false},
	}

	for _, c := range cases {
		name, err := ExpandTag(c.input, "")
		if c.valid {
			if err != nil {
				t.Errorf("ExpandTag(%q) returned error: %s", c.input, err)
			}
			if name != c.expected {
				t.Errorf("ExpandTag(%q) = %q, want %q", c.input, name, c.expected)
			}
		} else {
			if err == nil {
				t.Errorf("ExpandTag(%q) should return an error.", c.input)
			}
		}
	}
}

// TestDisplayProvidersJSON checks the provider list JSON output.
func TestDisplayProvidersJSON(t *testing.T) {
	setupTestProviders(t)

	out, err := DisplayProvidersJSON(false, "")
	if err != nil {
		t.Fatalf("DisplayProvidersJSON(false) returned error: %s", err)
	}

	expected := "[\"alpha\",\"beta\",\"gamma\"]\n"
	if out != expected {
		t.Errorf("DisplayProvidersJSON(false) = %q, want %q", out, expected)
	}
}

// TestDisplayProvidersJSONVerbose checks the verbose provider list JSON output.
func TestDisplayProvidersJSONVerbose(t *testing.T) {
	setupTestProviders(t)

	out, err := DisplayProvidersJSON(true, "")
	if err != nil {
		t.Fatalf("DisplayProvidersJSON(true) returned error: %s", err)
	}

	expected := `[
  {
    "provider": "alpha",
    "tags": [
      "code",
      "linux"
    ]
  },
  {
    "provider": "beta",
    "tags": [
      "code"
    ]
  },
  {
    "provider": "gamma",
    "tags": []
  }
]
`
	if out != expected {
		t.Errorf("DisplayProvidersJSON(true) = %q, want %q", out, expected)
	}
}

// TestDisplayTagsJSON checks the tag list JSON output.
func TestDisplayTagsJSON(t *testing.T) {
	setupTestProviders(t)

	out, err := DisplayTagsJSON(false, "")
	if err != nil {
		t.Fatalf("DisplayTagsJSON(false) returned error: %s", err)
	}

	expected := "[\"code\",\"linux\"]\n"
	if out != expected {
		t.Errorf("DisplayTagsJSON(false) = %q, want %q", out, expected)
	}
}

// TestDisplayTagsJSONVerbose checks the verbose tag list JSON output.
func TestDisplayTagsJSONVerbose(t *testing.T) {
	setupTestProviders(t)

	out, err := DisplayTagsJSON(true, "")
	if err != nil {
		t.Fatalf("DisplayTagsJSON(true) returned error: %s", err)
	}

	expected := `[
  {
    "tag": "code",
    "providers": [
      "alpha",
      "beta"
    ]
  },
  {
    "tag": "linux",
    "providers": [
      "alpha"
    ]
  }
]
`
	if out != expected {
		t.Errorf("DisplayTagsJSON(true) = %q, want %q", out, expected)
	}
}
