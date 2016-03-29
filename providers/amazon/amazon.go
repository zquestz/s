package amazon

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("amazon", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Amazon.
func (p *Provider) BuildURI(q string) string {
	switch providers.Region() {
	case "AU":
		return fmt.Sprintf("https://www.amazon.com.au/s/?keywords=%s", url.QueryEscape(q))
	case "BR":
		return fmt.Sprintf("https://www.amazon.com.br/s/?keywords=%s", url.QueryEscape(q))
	case "CA":
		return fmt.Sprintf("https://www.amazon.ca/s/?keywords=%s", url.QueryEscape(q))
	case "DE":
		return fmt.Sprintf("https://www.amazon.de/s/?keywords=%s", url.QueryEscape(q))
	case "ES":
		return fmt.Sprintf("https://www.amazon.es/s/?keywords=%s", url.QueryEscape(q))
	case "FR":
		return fmt.Sprintf("https://www.amazon.fr/s/?keywords=%s", url.QueryEscape(q))
	case "GB":
		return fmt.Sprintf("https://www.amazon.co.uk/s/?keywords=%s", url.QueryEscape(q))
	case "IN":
		return fmt.Sprintf("https://www.amazon.in/s/?keywords=%s", url.QueryEscape(q))
	case "IT":
		return fmt.Sprintf("https://www.amazon.it/s/?keywords=%s", url.QueryEscape(q))
	case "JP":
		return fmt.Sprintf("https://www.amazon.co.jp/s/?keywords=%s", url.QueryEscape(q))
	case "MX":
		return fmt.Sprintf("https://www.amazon.com.mx/s/?keywords=%s", url.QueryEscape(q))
	case "NL":
		return fmt.Sprintf("https://www.amazon.nl/s/?keywords=%s", url.QueryEscape(q))
	default:
		return fmt.Sprintf("https://www.amazon.com/s/?keywords=%s", url.QueryEscape(q))
	}
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"shopping"}
}
