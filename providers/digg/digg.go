package digg

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("digg", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for Digg.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://digg.com/search?q=%s", url.QueryEscape(q))
}
