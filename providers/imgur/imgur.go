package imgur

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("imgur", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for Imgur.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://imgur.com/search?q=%s", url.QueryEscape(q))
}
