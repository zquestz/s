package packagist

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("packagist", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for packagist.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://packagist.org/search/?q=%s", url.QueryEscape(q))
}
