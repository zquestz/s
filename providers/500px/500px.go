package fivehundredpx

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("500px", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for 8tracks.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://500px.com/search?q=%s&type=photos", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"photos"}
}
