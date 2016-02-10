package debianpkg

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("debianpkg", &Provider{})
}

// Provider adheres to the Provider interface.
type Provider struct {
}

// BuildURI generates the search URL for packages.debian.org.
func (p *DebianPKGProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://packages.debian.org/search?keywords=%s&searchon=names&suite=stable&section=all", url.QueryEscape(q))
}
