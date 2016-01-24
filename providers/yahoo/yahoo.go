package yahoo

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("yahoo", &YahooProvider{})
}

// YahooProvider adheres to the Provider interface.
type YahooProvider struct {
}

// BuildURI generates a search URL for Yahoo.
func (p *YahooProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://search.yahoo.com/search?p=%s", url.QueryEscape(q))
}
