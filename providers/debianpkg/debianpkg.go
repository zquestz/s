package debianpkg

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("debianpkg", &DebianPKGProvider{})
}

// DebianPKGProvider adheres to the Provider interface.
type DebianPKGProvider struct {
}

// BuildURI generates a search URL for WolframAlpha.
func (p *DebianPKGProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://packages.debian.org/search?keywords=%s&searchon=names&suite=stable&section=all", url.QueryEscape(q))
}
