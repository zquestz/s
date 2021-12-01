package idealo

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("idealo", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for IDEALO.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.idealo.de/preisvergleich/MainSearchProductCategory.html?q=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	switch providers.Language() {
	case "de":
		return []string{"shopping"}
	default:
		return []string{}
}
