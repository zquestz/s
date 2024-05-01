package flake8rules

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("flake8rules", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Flake8Rules.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.flake8rules.com/rules/%s.html", url.QueryEscape(strings.ToUpper(q)))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{}
}
