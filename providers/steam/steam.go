package steam

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("steam", &SteamProvider{})
}

// SteamProvider adheres to the Provider interface.
type SteamProvider struct {
}

// BuildURI generates a search URL for Steam.
func (p *SteamProvider) BuildURI(q string) string {
	return fmt.Sprintf(
		"http://store.steampowered.com/search/?term=%s", url.QueryEscape(q))
}
