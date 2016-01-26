package npm

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("npm", &NpmProvider{})
}

// NpmProvider adheres to the Provider interface.
type NpmProvider struct {
}

// BuildURI generates a search URL for npm.
func (p *NpmProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.npmjs.com/search?q=%s", url.QueryEscape(q))
}
