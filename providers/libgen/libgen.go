package libgen

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("libgen", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Libgen.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("http://gen.lib.rus.ec/search.php?req=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"education"}
}
