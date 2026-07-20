package arstechnica

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("arstechnica", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for arstechnica.
func (p *Provider) BuildURI(q string, _ string) string {
	return fmt.Sprintf("https://arstechnica.com/search/?ie=UTF-8&q=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags(_ string) []string {
	return []string{"tech-news"}
}
