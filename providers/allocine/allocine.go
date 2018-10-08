package allocine

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("allocine", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Allocine.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("http://www.allocine.fr/recherche/?q=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	switch providers.Language() {
	case "fr":
		return []string{"movies"}
	default:
		return []string{}
	}
}
