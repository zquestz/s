package freshports

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("freshports", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for freshports.org ports.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.freshports.org/search.php?query=go&search=%s&num=10&stype=name&method=match&deleted=excludedeleted&start=1&casesensitivity=caseinsensitive", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{}
}
