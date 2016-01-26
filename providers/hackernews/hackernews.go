package hackernews

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("hackernews", &Hackernews{})
}

// Hackernews adheres to the Provider interface.
type Hackernews struct {
}

// BuildURI generates a search URL for Hackernews.
func (p *Hackernews) BuildURI(q string) string {
	return fmt.Sprintf("https://hn.algolia.com/?q=%s", url.QueryEscape(q))
}
