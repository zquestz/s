package github

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("github", &GitHubProvider{})
}

// GitHubProvider adheres to the Provider interface.
type GitHubProvider struct {
}

// BuildURI generates a search URL for GitHub.
func (p *GitHubProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://github.com/search?utf8=✓&q=%s", url.QueryEscape(q))
}
