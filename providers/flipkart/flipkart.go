package flipkart

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("flipkart", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for Flipkart.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("http://www.flipkart.com/search?q=%s", url.QueryEscape(q))
}
