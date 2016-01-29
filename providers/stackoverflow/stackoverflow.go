package stackoverflow

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("stackoverflow", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for Stack Overflow.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://stackoverflow.com/search?q=%s", url.QueryEscape(q))
}
