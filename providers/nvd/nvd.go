package nvd

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("nvd", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for the National Vulnerability Database.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://web.nvd.nist.gov/view/vuln/search-results?query=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{}
}
