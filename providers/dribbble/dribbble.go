package dribbble

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("dribbble", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for Dribbble.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://dribbble.com/search?q=%s", url.QueryEscape(q))
}
