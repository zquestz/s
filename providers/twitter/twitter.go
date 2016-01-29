package twitter

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("twitter", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for Twitter.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://twitter.com/search?q=%s", url.QueryEscape(q))
}
