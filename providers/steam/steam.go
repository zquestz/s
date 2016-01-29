package steam

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("steam", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for Steam.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://store.steampowered.com/search/?term=%s", url.QueryEscape(q))
}
