package caskroom

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("caskroom", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Caskroom.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("caskroom.io/search?q=%s, url.QueryEscape(q))
}
