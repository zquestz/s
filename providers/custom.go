package providers

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var alphanums = regexp.MustCompile("^[a-zA-Z0-9_]*$")

// CustomProvider is used for Config based providers.
type CustomProvider struct {
	Name    string   `json:"name"`
	URL     string   `json:"url"`
	TagList []string `json:"tags"`
}

// BuildURI builds the URI for custom providers.
func (c *CustomProvider) BuildURI(q string) string {
	return fmt.Sprintf(c.URL, url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (c *CustomProvider) Tags() []string {
	return c.TagList
}

// Valid checks if the custom provider is setup correctly.
func (c *CustomProvider) Valid() error {
	// Validate name is only alphanums.
	if !alphanums.Match([]byte(c.Name)) {
		return fmt.Errorf("name must be alphanumeric")
	}

	u, err := url.Parse(c.URL)
	if err != nil {
		return err
	}

	c.URL = u.String()

	// Make sure query token is present
	hasToken := strings.Contains(c.URL, "%s")
	if !hasToken {
		return fmt.Errorf("token %q required", "%s")
	}

	// Validate scheme is set. Don't limit to http.
	if u.Scheme == "" {
		return fmt.Errorf("url scheme required")
	}

	return nil
}
