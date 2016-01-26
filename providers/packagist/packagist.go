package packagist

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("packagist", &PackagistProvider{})
}

// PackagistProvider adheres to the Provider interface.
type PackagistProvider struct {
}

// BuildURI generates a search URL for packagist.
func (p *PackagistProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://packagist.org/search/?q=%s", url.QueryEscape(q))
}
