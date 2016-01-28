package dribbble

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("dribbble", &DribbbleProvider{})
}

//DribbbleProvider adheres to the Provider interface.
type DribbbleProvider struct {
}

//BuildURI generates a search URL for Dribbble.
func (p *DribbbleProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://dribbble.com/search?q=%s", url.QueryEscape(q))
}
