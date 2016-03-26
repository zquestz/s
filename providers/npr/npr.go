package npr

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("npr", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for NPR.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("http://www.npr.org/templates/search/index.php?searchinput=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"news"}
}
