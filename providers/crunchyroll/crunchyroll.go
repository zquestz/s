package crunchyroll

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("crunchyroll", &CrunchyrollProvider{})
}

// CrunchyrollProvider adheres to the Provider interface.
type CrunchyrollProvider struct {
}

// BuildURI generates a search URL for Crunchyroll.
func (p *CrunchyrollProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.crunchyroll.com/search?q=%s", url.QueryEscape(q))
}
