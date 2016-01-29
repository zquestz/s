package godoc

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("godoc", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for GoDoc.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://godoc.org/?q=%s", url.QueryEscape(q))
}
