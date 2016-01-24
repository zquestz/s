package duckduckgo

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("duckduckgo", &DuckProvider{})
}

// DuckProvider adheres to the Provider interface.
type DuckProvider struct {
}

// BuildURI generates a search URL for DuckDuckGo.
func (p *DuckProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://duckduckgo.com/?q=%s", url.QueryEscape(q))
}
