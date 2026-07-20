package golang

import (
	"fmt"
	"net/url"
	"slices"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("flathub", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Flathub.
func (p *Provider) BuildURI(q string, locale string) string {
	flathubLanguages := []string{
		"az",
		"br",
		"en-GB",
		"ca",
		"cs",
		"da",
		"de",
		"et",
		"en",
		"es",
		"eo",
		"eu",
		"fil",
		"fr",
		"ga",
		"gl",
		"hr",
		"id",
		"ia",
		"it",
		"kw",
		"lt",
		"hu",
		"nl",
		"nb-NO",
		"oc",
		"pl",
		"pt",
		"pt-BR",
		"ro",
		"sq",
		"fi",
		"sv",
		"kab",
		"vi",
		"tr",
		"el",
		"be",
		"bg",
		"ru",
		"uk",
		"hy",
		"he",
		"ar",
		"fa",
		"ckb",
		"hi",
		"bn",
		"pa",
		"ta",
		"si",
		"ko",
		"ja",
		"zh-Hans",
		"zh-Hant",
	}

	userLanguage := providers.Language(locale)

	if slices.Contains(flathubLanguages, userLanguage) {
		return fmt.Sprintf("https://flathub.org/%s/apps/search?q=%s", userLanguage, url.QueryEscape(q))
	}

	return fmt.Sprintf("https://flathub.org/en/apps/search?q=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags(_ string) []string {
	return []string{}
}
