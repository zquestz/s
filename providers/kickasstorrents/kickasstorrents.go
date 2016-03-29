package kickasstorrents

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("kickasstorrents", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for KickAss Torrents.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://kat.cr/usearch/%s/", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"torrents"}
}
