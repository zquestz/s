package quora

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("quora", &QuoraProvider{})
}

// QuoraProvider adheres to the Provider interface.
type QuoraProvider struct {
}

// BuildURI generates a search URL for Quora.
func (p *QuoraProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.quora.com/search?q=%s", url.QueryEscape(q))
}
