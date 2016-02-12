package rottentomatoes

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("rottentomatoes", &RottenTomatoesProvider{})
}

// RottenTomatoesProvider adheres to the Provider interface.
type RottenTomatoesProvider struct {
}

// BuildURI generates a search URL for RottenTomatoes.
func (p *RottenTomatoesProvider) BuildURI(q string) string {
	return fmt.Sprintf("http://www.rottentomatoes.com/search/?search=%s", url.QueryEscape(q))
}
