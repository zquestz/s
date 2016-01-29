package rubygems

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("rubygems", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for RubyGems.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://rubygems.org/search?utf8=✓&query=%s", url.QueryEscape(q))
}
