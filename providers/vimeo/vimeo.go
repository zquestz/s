package vimeo

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("vimeo", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for Vimeo.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://vimeo.com/search?q=%s", url.QueryEscape(q))
}
