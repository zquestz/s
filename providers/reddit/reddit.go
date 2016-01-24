package reddit

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("reddit", &RedditProvider{})
}

// RedditProvider adheres to the Provider interface.
type RedditProvider struct {
}

// BuildURI generates a search URL for Reddit.
func (p *RedditProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.reddit.com/search?q=%s", url.QueryEscape(q))
}
