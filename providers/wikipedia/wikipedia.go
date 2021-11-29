package wikipedia

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("wikipedia", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Wikipedia.
func (p *Provider) BuildURI(q string) string {
	switch providers.Region() {
	case "BR":
		return fmt.Sprintf("https://br.wikipedia.org/wiki/Special:Search?search=%s", url.QueryEscape(q))
	case "CA":
		return fmt.Sprintf("https://ca.wikipedia.org/wiki/Special:Search?search=%s", url.QueryEscape(q))
	case "DE":
		return fmt.Sprintf("https://de.wikipedia.org/wiki/Special:Search?search=%s", url.QueryEscape(q))
	case "ES":
		return fmt.Sprintf("https://es.wikipedia.org/wiki/Special:Search?search=%s", url.QueryEscape(q))
	case "FR":
		return fmt.Sprintf("https://fr.wikipedia.org/wiki/Special:Search?search=%s", url.QueryEscape(q))
	case "IT":
		return fmt.Sprintf("https://it.wikipedia.org/wiki/Special:Search?search=%s", url.QueryEscape(q))
	case "JP":
		return fmt.Sprintf("https://jp.wikipedia.org/wiki/Special:Search?search=%s", url.QueryEscape(q))
	case "NL":
		return fmt.Sprintf("https://nl.wikipedia.org/wiki/Special:Search?search=%s", url.QueryEscape(q))
	default:
		return fmt.Sprintf("https://en.wikipedia.org/wiki/Special:Search?search=%s", url.QueryEscape(q))
	}
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"education"}
}
