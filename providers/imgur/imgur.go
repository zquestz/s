package imgur

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("imgur", &Imgur{})
}

// Imgur adheres to the Provider interface.
type Imgur struct {
}

// BuildURI generates a search URL for Imgur.
func (p *Imgur) BuildURI(q string) string {
	return fmt.Sprintf("https://imgur.com/search?q=%s", url.QueryEscape(q))
}
