package dumpert

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("dumpert", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for dumpert.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf(
		"https://www.dumpert.nl/zoek/%s",
		strings.Replace(url.QueryEscape(q), "+", "%20", -1),
	)
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{}
}
