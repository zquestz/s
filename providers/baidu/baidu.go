package baidu

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("baidu", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for Baidu.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.baidu.com/s?wd=%s", url.QueryEscape(q))
}
