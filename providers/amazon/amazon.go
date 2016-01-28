package amazon

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("amazon", &AmazonProvider{})
}

// AmazonProvider adheres to the Provider interface.
type AmazonProvider struct {
}

// BuildURI generates a search URL for Amazon.
func (p *AmazonProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.amazon.com/s/?keywords=%s", url.QueryEscape(q))
}
