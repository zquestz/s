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

	// Make sure query token is present
	hasToken := strings.Contains(c.URL, "%s")
	if !hasToken {
		return fmt.Errorf("token %q required", "%s")
	}

	// Substitute %s before parsing since url.Parse()
	// fails if the path contains %s.
	u, err := url.Parse(fmt.Sprintf(c.URL, "search_term"))
	if err != nil {
		return err
	}

	// Validate scheme is set. Don't limit to http.
	if u.Scheme == "" {
		return fmt.Errorf("url scheme required")
	}

	return nil
}
