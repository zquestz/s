package hackernews

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("hackernews", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for Hackernews.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://hn.algolia.com/?q=%s", url.QueryEscape(q))
}
