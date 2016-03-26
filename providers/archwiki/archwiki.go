package archwiki

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("archwiki", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for the ArchWiki.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://wiki.archlinux.org/index.php?search=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"arch"}
}
