package kaufda

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("kaufda", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for kaufda.de.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.kaufda.de/webapp/?query=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	switch providers.Language() {
	case "de":
		return []string{"shopping"}
	default:
		return []string{}
	}
}
