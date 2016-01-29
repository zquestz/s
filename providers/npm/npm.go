package npm

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("npm", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for npm.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.npmjs.com/search?q=%s", url.QueryEscape(q))
}
