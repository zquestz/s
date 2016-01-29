package gist

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("gist", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for Gist.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://gist.github.com/search?q=%s", url.QueryEscape(q))
}
