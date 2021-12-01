package ardmediathek

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("ardmediathek", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for ARD MEDIATHEK.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.ardmediathek.de/suche/%s/", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	switch providers.Language() {
	case "de":
		return []string{"tv"}
	default:
		return []string{}
	}
}
