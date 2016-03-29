package wolframalpha

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("wolframalpha", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for WolframAlpha.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.wolframalpha.com/input/?i=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{}
}
