package github

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("github", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for GitHub.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://github.com/search?utf8=✓&q=%s", url.QueryEscape(q))
}
