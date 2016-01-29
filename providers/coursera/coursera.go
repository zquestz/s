package coursera

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("coursera", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for Coursera.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.coursera.org/courses/?query=%s", url.QueryEscape(q))
}
