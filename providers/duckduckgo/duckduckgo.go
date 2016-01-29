package duckduckgo

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("duckduckgo", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for DuckDuckGo.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://duckduckgo.com/?q=%s", url.QueryEscape(q))
}
