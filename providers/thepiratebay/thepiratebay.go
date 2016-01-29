package thepiratebay

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("thepiratebay", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for ThePirateBay.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://thepiratebay.se/search/%s", url.QueryEscape(q))
}
