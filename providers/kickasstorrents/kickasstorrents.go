package kickasstorrents

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("kickasstorrents", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for KickAss Torrents.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://kat.cr/usearch/%s/", url.QueryEscape(q))
}
