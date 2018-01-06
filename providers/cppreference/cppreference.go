package cppreference

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("cppreference", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for cppreference.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("http://www.cppreference.com/mwiki/index.php?search=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{}
}
