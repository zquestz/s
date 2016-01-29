package yahoo

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("yahoo", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for Yahoo.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://search.yahoo.com/search?p=%s", url.QueryEscape(q))
}
