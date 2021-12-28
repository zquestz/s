package openbsdman

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("openbsdman", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for the OpenBSD online man pages dictionary.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://man.openbsd.org/%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{}
}
