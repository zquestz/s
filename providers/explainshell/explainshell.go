package explainshell

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

const (
	providerName = "explainshell"
	providerURL  = "https://explainshell.com"
)

func init() {
	providers.AddProvider(providerName, &Provider{})
}

// Provider merely implements the Provider interface
type Provider struct{}

// BuildURI generates a search URL for Explain Shell
func (p *Provider) BuildURI(q string) string {
	searchParam := url.QueryEscape(q)
	return fmt.Sprintf("%s/explain?cmd=%s", providerURL, searchParam)
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"code"}
}
