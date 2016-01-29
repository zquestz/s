package youtube

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("youtube", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for YouTube.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.youtube.com/results?search_query=%s", url.QueryEscape(q))
}
