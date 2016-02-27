package providers

import "testing"

var (
	validName = "example"
	validURL  = "http://example.com/?q=%s"
)

// TestValid checks validation errors on custom providers.
func TestValid(t *testing.T) {
	cases := []struct {
		name  string
		url   string
		valid bool
	}{
		{validName, validURL, true},
		{"_", validURL, true},
		{"-", validURL, false},
		{"with whitespace", validURL, false},
		{validName, "noscheme.com?q=%s", false},
		{validName, "http://notoken.com", false},
		{validName, "invalid uri", false},
	}

	for _, c := range cases {
		p := CustomProvider{
			Name: c.name,
			URL:  c.url,
		}

		err := p.Valid()
		if c.valid {
			if err != nil {
				t.Errorf("Provider %#v should be valid but returned error: %s", p, err)
			}
		} else {
			if err == nil {
				t.Errorf("Provider %#v should be invalid but didn't return an error.", p)
			}
		}
	}
}
