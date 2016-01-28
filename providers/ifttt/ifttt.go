package ifttt

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("ifttt", &IFTTTProvider{})
}

// IFTTTProvider adheres to the Provider interface.
type IFTTTProvider struct {
}

// BuildURI generates a search URL for IFTTT.
func (p *IFTTTProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.ifttt.com/recipes/search?q=%s", url.QueryEscape(q))
}
