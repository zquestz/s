package ultimateguitar

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("ultimateguitar", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for ultimate-guitar.com.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.ultimate-guitar.com/search.php?search_type=title&value=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{}
}
