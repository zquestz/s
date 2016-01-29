package golang

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("go", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for Go.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://golang.org/search?q=%s", url.QueryEscape(q))
}
