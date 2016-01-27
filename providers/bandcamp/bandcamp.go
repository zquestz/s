package bandcamp

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("bandcamp", &BandcampProvider{})
}

// BandcampProvider adheres to the Provider interface.
type BandcampProvider struct {
}

// BuildURI generates a search URL for Bandcamp.
func (p *BandcampProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://bandcamp.com/search?q=%s", url.QueryEscape(q))
}
