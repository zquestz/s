package providers

import "testing"

// setupJSONProviders installs a fixed set of providers
// and restores the original providers after the test.
func setupJSONProviders(t *testing.T) {
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

// TestDisplayProvidersJSON checks the provider list JSON output.
func TestDisplayProvidersJSON(t *testing.T) {
	setupJSONProviders(t)

	out, err := DisplayProvidersJSON(false)
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
	setupJSONProviders(t)

	out, err := DisplayProvidersJSON(true)
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
	setupJSONProviders(t)

	out, err := DisplayTagsJSON(false)
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
	setupJSONProviders(t)

	out, err := DisplayTagsJSON(true)
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
