package instagram

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("instagram", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Php.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.instagram.com/explore/tags/%s/", url.QueryEscape(q))
}
