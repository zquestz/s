package yandex

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("yandex", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for Yandex.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.yandex.com/search/?text=%s", url.QueryEscape(q))
}
