package google

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("google", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Google.
func (p *Provider) BuildURI(q string) string {
	switch providers.Region() {
	case "CA":
		return fmt.Sprintf("https://www.google.ca/search?q=%s", url.QueryEscape(q))
	case "DE":
		return fmt.Sprintf("https://www.google.de/search?q=%s", url.QueryEscape(q))
	case "ES":
		return fmt.Sprintf("https://www.google.es/search?q=%s", url.QueryEscape(q))
	case "FR":
		return fmt.Sprintf("https://www.google.fr/search?q=%s", url.QueryEscape(q))
	case "IN":
		return fmt.Sprintf("https://www.google.co.in/search?q=%s", url.QueryEscape(q))
	case "IT":
		return fmt.Sprintf("https://www.google.it/search?q=%s", url.QueryEscape(q))
	case "JP":
		return fmt.Sprintf("https://www.google.co.jp/search?q=%s", url.QueryEscape(q))
	case "MX":
		return fmt.Sprintf("https://www.google.com.mx/search?q=%s", url.QueryEscape(q))
	case "NL":
		return fmt.Sprintf("https://www.google.nl/search?q=%s", url.QueryEscape(q))
	default:
		return fmt.Sprintf("https://www.google.com/search?q=%s", url.QueryEscape(q))
	}
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"search"}
}
