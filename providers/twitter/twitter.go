package twitter

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("twitter", &TwitterProvider{})
}

// TwitterProvider adheres to the Provider interface.
type TwitterProvider struct {
}

// BuildURI generates a search URL for Twitter.
func (p *TwitterProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://twitter.com/search?q=%s", url.QueryEscape(q))
}
