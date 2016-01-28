package flipkart

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("flipkart", &FlipkartProvider{})
}

// FlipkartProvider adheres to the Provider interface.
type FlipkartProvider struct {
}

// BuildURI generates a search URL for Flipkart.
func (p *FlipkartProvider) BuildURI(q string) string {
	return fmt.Sprintf("http://www.flipkart.com/search?q=%s", url.QueryEscape(q))
}
