package npmsearch

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("npmsearch", &NpmSearchProvider{})
}

// NpmSearchProvider adheres to the Provider interface.
type NpmSearchProvider struct {
}

// BuildURI generates a search URL for npmsearch.
func (p *NpmSearchProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://npmsearch.com/?q=%s", url.QueryEscape(q))
}
