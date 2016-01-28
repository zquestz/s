package amazonuk

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("amazonuk", &AmazonUkProvider{})
}

// AmazonProvider adheres to the Provider interface.
type AmazonUkProvider struct {
}

// BuildURI generates a search URL for Amazon.
func (p *AmazonUkProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.amazon.co.uk/s/?keywords=%s", url.QueryEscape(q))
}
