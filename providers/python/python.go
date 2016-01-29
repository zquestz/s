package python

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("python", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for Python.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.python.org/search/?q=%s", url.QueryEscape(q))
}
