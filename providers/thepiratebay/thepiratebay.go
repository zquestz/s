package thepiratebay

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("thepiratebay", &ThePirateBayProvider{})
}

// QuoraProvider adheres to the Provider interface.
type ThePirateBayProvider struct {
}

// BuildURI generates a search URL for ThePirateBay.
func (p *ThePirateBayProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://thepiratebay.se/search/%s", url.QueryEscape(q))
}
