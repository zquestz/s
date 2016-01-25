package pinterest

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("pinterest", &PinterestProvider{})
}

// PinterestProvider adheres to the Provider interface.
type PinterestProvider struct {
}

// BuildURI generates a search URL for Pinterest.
func (p *PinterestProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.pinterest.com/search/pins/?q=%s", url.QueryEscape(q))
}
