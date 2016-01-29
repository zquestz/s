package twitchtv

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("twitchtv", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for Twitchtv.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.twitch.tv/search?query=%s", url.QueryEscape(q))
}
