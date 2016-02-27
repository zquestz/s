package pypi

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("pypi", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct {
}

// BuildURI generates a search URL for PyPI.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://pypi.python.org/pypi?%3Aaction=search&term=%s&submit=search", url.QueryEscape(q))
}