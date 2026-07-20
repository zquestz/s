package golang

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("unsplash", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Unsplash.
func (p *Provider) BuildURI(q string, locale string) string {
	unsplashExtraLanguages := map[string]string{
		"de":    "https://unsplash.com/de/s/fotos/%s",
		"es":    "https://unsplash.com/es/s/fotos/%s",
		"fr":    "https://unsplash.com/fr/s/photos/%s",
		"id":    "https://unsplash.com/id/s/foto/%s",
		"it":    "https://unsplash.com/it/s/foto/%s",
		"ja":    "https://unsplash.com/ja/s/写真/%s",
		"ko":    "https://unsplash.com/ko/s/사진/%s",
		"pt-br": "https://unsplash.com/pt-br/s/fotografias/%s",
	}

	if languageURLFormatString, exists := unsplashExtraLanguages[providers.Language(locale)]; exists {
		return fmt.Sprintf(languageURLFormatString, url.QueryEscape(q))
	}

	return fmt.Sprintf("https://unsplash.com/s/photos/%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags(_ string) []string {
	return []string{"photos"}
}
