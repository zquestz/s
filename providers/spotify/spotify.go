package spotify

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("spotify", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for spotify.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://play.spotify.com/search/%s", url.QueryEscape(q))
}
