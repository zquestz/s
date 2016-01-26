package packagistsearch

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("packagistsearch", &PackagistSearchProvider{})
}

// PackagistSearchProvider adheres to the Provider interface.
type PackagistSearchProvider struct {
}

// BuildURI generates a search URL for packagistsearch.
func (p *PackagistSearchProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://packagist.org/search/?q=%s", url.QueryEscape(q))
}
