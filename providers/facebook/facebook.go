package facebook

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("facebook", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for facebook.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.facebook.com/search/top/?q=%s", url.QueryEscape(q))
}
