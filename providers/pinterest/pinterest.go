package pinterest

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("pinterest", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for Pinterest.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.pinterest.com/search/pins/?q=%s", url.QueryEscape(q))
}
