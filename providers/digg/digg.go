package digg

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("digg", &DiggProvider{})
}

// DiggProvider adheres to the Provider interface.
type DiggProvider struct {
}

// BuildURI generates a search URL for Digg.
func (p *DiggProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://digg.com/search?q=%s", url.QueryEscape(q))
}
