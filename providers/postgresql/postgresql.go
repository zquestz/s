package postgresql

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("postgresql", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Google.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("http://www.postgresql.org/search/?q=%s&a=1&submit=Search", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"search"}
}
