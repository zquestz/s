package reddit

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("reddit", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Reddit.
// If q starts with "/r/subreddit query", treat as subreddit and do restrict
// search on the subreddit. The exception is when q = "/r/subreddit" where no
// query is provided, treat it as regular reddit search.
func (p *Provider) BuildURI(q string) string {
	qs := strings.Split(q, " ")
	if strings.HasPrefix(qs[0], "/r/") && len(qs) > 1 {
		return fmt.Sprintf("https://www.reddit.com%s/search?q=%s&restrict_sr=on",
			qs[0], url.QueryEscape(strings.Join(qs[1:], " ")))
	}
	return fmt.Sprintf("https://www.reddit.com/search?q=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"forums"}
}
