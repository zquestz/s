package reddit

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("reddit", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Reddit.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.reddit.com/search?q=%s", url.QueryEscape(q))
}
