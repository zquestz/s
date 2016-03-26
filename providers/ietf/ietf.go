package ietf

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("ietf", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for IETF.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://datatracker.ietf.org/doc/search/"+
		"?name=%s&rfcs=on&activedrafts=on", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{}
}
