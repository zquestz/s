package google

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("google", &GoogleProvider{})
}

// GoogleProvider adheres to the Provider interface.
type GoogleProvider struct {
}

// BuildURI generates a search URL for Google.
func (p *GoogleProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.google.com/search?q=%s", url.QueryEscape(q))
}
