package yandex

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("yandex", &YandexProvider{})
}

// YandexProvider adheres to the Provider interface.
type YandexProvider struct {
}

// BuildURI generates a search URL for Google.
func (p *YandexProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.yandex.com/search/?text=%s", url.QueryEscape(q))
}
