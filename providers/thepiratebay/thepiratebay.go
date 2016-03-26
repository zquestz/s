package thepiratebay

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("thepiratebay", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for ThePirateBay.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://thepiratebay.se/search/%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"torrents"}
}
