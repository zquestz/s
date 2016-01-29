package npmsearch

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("npmsearch", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for npmsearch.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://npmsearch.com/?q=%s", url.QueryEscape(q))
}
