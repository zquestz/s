package crates

import (
	"fmt"
	"net/url"

	"github.com/anykao/s/providers"
)

func init() {
	providers.AddProvider("crates", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for crates.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://crates.io/search?q=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{}
}
