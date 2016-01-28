package dumpert

import (
	"fmt"
	"net/url"
	"os/exec"

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
	_, err := exec.LookPath("spotify")
	if err != nil {
		return fmt.Sprintf("https://play.spotify.com/search/%s", url.QueryEscape(q))
	}
	
	return fmt.Sprintf("spotify:search:%s", url.QueryEscape(q))
}
