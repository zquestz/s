package rubygems

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("rubygems", &RubyGemsProvider{})
}

// RubyGemsProvider adheres to the Provider interface.
type RubyGemsProvider struct {
}

// BuildURI generates a search URL for RubyGems.
func (p *RubyGemsProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://rubygems.org/search?utf8=✓&query=%s", url.QueryEscape(q))
}
