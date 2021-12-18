package freebsdman

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("freebsdman", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for the FreeBSD online man pages dictionary.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.freebsd.org/cgi/man.cgi?query=%s&apropos=0&sektion=0&arch=default&format=html", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{}
}
