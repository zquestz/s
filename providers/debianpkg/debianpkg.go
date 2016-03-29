package debianpkg

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("debianpkg", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for debian.org packages.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://packages.debian.org/search?keywords=%s&searchon=names&suite=stable&section=all", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{}
}
