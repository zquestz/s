package steam

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("steam", &Steam{})
}

// Steam adheres to the Provider interface.
type Steam struct {
}

// BuildURI generates a search URL for Steam.
func (p *Steam) BuildURI(q string) string {
	return fmt.Sprintf("http://store.steampowered.com/search/?term=%s", url.QueryEscape(q))
}
