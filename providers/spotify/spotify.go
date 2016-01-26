package dumpert

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("spotify", &SpotifyProvider{})
}

// SpotifyProvider adheres to the Provider interface.
type SpotifyProvider struct {
}

// BuildURI generates a search URL for spotify.
func (p *SpotifyProvider) BuildURI(q string) string {
	// https://play.spotify.com/search/%s
	return fmt.Sprintf("spotify:search:%s", url.QueryEscape(q))
}
