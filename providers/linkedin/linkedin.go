package linkedin

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("linkedin", &LinkedinProvider{})
}

// LinkedinProvider adheres to the Provider interface.
type LinkedinProvider struct {
}

// BuildURI generates a search URL for Linkedin Provider.
func (p *LinkedinProvider) BuildURI(q string) string {
	return fmt.Sprintf(
		"https://www.linkedin.com/vsearch/f?type=all&keywords=%s&search=Search",
		url.QueryEscape(q))
}
