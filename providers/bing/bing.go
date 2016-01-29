package bing

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("bing", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for Bing.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.bing.com/search?q=%s", url.QueryEscape(q))
}
