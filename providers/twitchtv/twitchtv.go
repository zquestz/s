package twitchtv

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("twitchtv", &Twitchtv{})
}

// Twitchtv adheres to the Provider interface.
type Twitchtv struct {
}

// BuildURI generates a search URL for Twitchtv.
func (p *Twitchtv) BuildURI(q string) string {
	return fmt.Sprintf("http://www.twitch.tv/search?query=%s", url.QueryEscape(q))
}
