package baidu

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("baidu", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Baidu.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.baidu.com/s?wd=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	switch providers.Language() {
	case "zh":
		return []string{"search"}
	default:
		return []string{}
	}
}
