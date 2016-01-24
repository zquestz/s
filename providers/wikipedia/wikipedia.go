package wikipedia

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("wikipedia", &WikiProvider{})
}

// WikiProvider adheres to the Provider interface.
type WikiProvider struct {
}

// BuildURI generates a search URL for Wikipedia.
func (p *WikiProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://en.wikipedia.org/wiki/Special:Search?search=%s", url.QueryEscape(q))
}
