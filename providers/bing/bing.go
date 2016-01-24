package bing

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("bing", &BingProvider{})
}

// BingProvider adheres to the Provider interface.
type BingProvider struct {
}

// BuildURI generates a search URL for Bing.
func (p *BingProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.bing.com/search?q=%s", url.QueryEscape(q))
}
