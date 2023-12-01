package flake8rules

import (
	"strings"
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("flake8rules", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for ProtonDB.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.flake8rules.com/rules/%s.html", url.QueryEscape(strings.ToUpper(q)))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{}
}
