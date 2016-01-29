package mdn

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("mdn", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for MDN.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://developer.mozilla.org/en-US/search?q=%s", url.QueryEscape(q))
}
