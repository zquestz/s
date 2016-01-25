package gist

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("gist", &GistProvider{})
}

// GistProvider adheres to the Provider interface.
type GistProvider struct {
}

// BuildURI generates a search URL for Gist.
func (p *GistProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://gist.github.com/search?q=%s", url.QueryEscape(q))
}
