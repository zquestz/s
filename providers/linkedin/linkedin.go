package linkedin

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("linkedin", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for Linkedin.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf(
		"https://www.linkedin.com/vsearch/f?type=all&keywords=%s&search=Search",
		url.QueryEscape(q))
}
