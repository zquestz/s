package arxiv

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("arxiv", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for arxiv.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://arxiv.org/search/?query=%s&searchtype=all", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"education"}
}
