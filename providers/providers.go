package providers

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/zquestz/s/launcher"
	"golang.org/x/text/language"
)

const (
	defaultLanguage = "en"
	defaultRegion   = "US"
)

var (
	enableBlacklist = false
	enableWhitelist = false
	clientLocale    = ""
	blacklist       = make(map[string]any)
	whitelist       = make(map[string]any)
)

// Provider interface provides a way to build the URI
// for each provider.
type Provider interface {
	BuildURI(string) string
	Tags() []string
}

// Providers tracks loaded providers.
var Providers map[string]Provider

func init() {
	Providers = make(map[string]Provider)
}

// AddProvider should be called within your provider's init() func.
// This will register the provider so it can be used.
func AddProvider(name string, provider Provider) {
	Providers[name] = provider
}

// SetBlacklist filters out unneeded providers.
func SetBlacklist(b []string) {
	for _, v := range b {
		blacklist[v] = nil
	}

	if len(blacklist) > 0 {
		enableBlacklist = true
	}
}

// SetWhitelist sets an exact list of supported providers.
func SetWhitelist(w []string) {
	for _, v := range w {
		whitelist[v] = nil
	}

	if len(whitelist) > 0 {
		enableWhitelist = true
	}
}

// Search builds a search URL and opens it in your browser.
func Search(binary string, p string, t string, q string, userProvider bool, outputOnly bool, verbose bool) error {
	prov, err := ExpandProvider(p)
	if err != nil {
		return err
	}

	builders := []Provider{}

	if t == "" || userProvider {
		builders = append(builders, Providers[prov])
	}

	if t != "" {
		tag, err := ExpandTag(t)
		if err != nil {
			return err
		}

		builders = append(builders, GetProvidersByTag(tag)...)
	}

	var success bool

	for _, builder := range builders {
		if builder != nil {
			url := builder.BuildURI(q)

			if verbose || outputOnly {
				fmt.Printf("%s\n", url)
			}

			if !outputOnly {
				err = launcher.OpenURI(binary, url)
				if err != nil {
					return err
				}
			}

			success = true
		}
	}

	if !success {
		return fmt.Errorf("no providers found for tag %q", t)
	}

	return nil
}

// GetProvidersByTag gets all providers with a specific tag.
func GetProvidersByTag(tag string) []Provider {
	provs := []Provider{}

	for _, providerName := range ProviderNames(false) {
		provider := Providers[providerName]
		for _, providerTag := range provider.Tags() {
			if providerTag == tag {
				provs = append(provs, provider)
			}
		}
	}

	return provs
}

// DisplayProviders displays all the loaded providers.
func DisplayProviders(verbose bool) string {
	names := ProviderNames(verbose)

	return fmt.Sprintf("%s\n", strings.Join(names, "\n"))
}

// DisplayTags displays all the available tags.
func DisplayTags(verbose bool) string {
	names := TagNames(verbose)

	return fmt.Sprintf("%s\n", strings.Join(names, "\n"))
}

// providerInfo is the JSON representation of a provider.
type providerInfo struct {
	Provider string   `json:"provider"`
	Tags     []string `json:"tags"`
}

// tagInfo is the JSON representation of a tag.
type tagInfo struct {
	Tag       string   `json:"tag"`
	Providers []string `json:"providers"`
}

// DisplayProvidersJSON returns all the loaded providers as JSON.
func DisplayProvidersJSON(verbose bool) (string, error) {
	if !verbose {
		data, err := json.Marshal(ProviderNames(false))
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s\n", data), nil
	}

	infos := []providerInfo{}

	for _, name := range ProviderNames(false) {
		tags := Providers[name].Tags()
		if tags == nil {
			tags = []string{}
		}
		sort.Strings(tags)

		infos = append(infos, providerInfo{Provider: name, Tags: tags})
	}

	data, err := json.MarshalIndent(infos, "", "  ")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s\n", data), nil
}

// DisplayTagsJSON returns all the available tags as JSON.
func DisplayTagsJSON(verbose bool) (string, error) {
	if !verbose {
		data, err := json.Marshal(TagNames(false))
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s\n", data), nil
	}

	tags := make(map[string][]string)

	for _, name := range ProviderNames(false) {
		for _, t := range Providers[name].Tags() {
			tags[t] = append(tags[t], name)
		}
	}

	infos := []tagInfo{}

	for _, tag := range TagNames(false) {
		infos = append(infos, tagInfo{Tag: tag, Providers: tags[tag]})
	}

	data, err := json.MarshalIndent(infos, "", "  ")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s\n", data), nil
}

// ExpandProvider expands the passed in provider to the full value.
func ExpandProvider(provider string) (string, error) {
	names := ProviderNames(false)
	r := regexp.MustCompile(`^` + provider)

	validProviders := []string{}
	for _, n := range names {
		// Exact match returns immediately.
		if n == provider {
			return n, nil
		}

		if r.Match([]byte(n)) {
			validProviders = append(validProviders, n)
		}
	}

	switch len(validProviders) {
	case 0:
		return "", fmt.Errorf("no provider found for %q", provider)
	case 1:
		return validProviders[0], nil
	default:
		return "", fmt.Errorf("multiple providers matched %q: %v", provider, validProviders)
	}
}

// ExpandTag expands the passed in tag to the full value.
func ExpandTag(tag string) (string, error) {
	names := TagNames(false)
	r := regexp.MustCompile(`^` + tag)

	validTags := []string{}
	for _, n := range names {
		// Exact match returns immediately.
		if n == tag {
			return n, nil
		}

		if r.Match([]byte(n)) {
			validTags = append(validTags, n)
		}
	}

	switch len(validTags) {
	case 0:
		return "", fmt.Errorf("no tag found for %q", tag)
	case 1:
		return validTags[0], nil
	default:
		return "", fmt.Errorf("multiple tags matched %q: %v", tag, validTags)
	}
}

// allowedProvider checks a provider name against
// both the whitelist and the blacklist.
func allowedProvider(name string) bool {
	if enableWhitelist {
		_, ok := whitelist[name]
		if !ok {
			return false
		}
	}

	if enableBlacklist {
		_, ok := blacklist[name]
		if ok {
			return false
		}
	}

	return true
}

// ProviderNames returns a sorted slice of provider names
// applying both the whitelist and then the blacklist.
func ProviderNames(verbose bool) []string {
	names := []string{}

	for key, p := range Providers {
		if !allowedProvider(key) {
			continue
		}

		if verbose {
			tags := p.Tags()

			if len(tags) > 0 {
				sort.Strings(tags)
				names = append(names, fmt.Sprintf("%s (%s)", key, strings.Join(tags, ", ")))
				continue
			}
		}

		names = append(names, key)
	}

	sort.Strings(names)
	return names
}

// TagNames returns a deduped and sorted slice of tag names.
func TagNames(verbose bool) []string {
	tags := make(map[string][]string)

	for name, p := range Providers {
		if !allowedProvider(name) {
			continue
		}

		for _, t := range p.Tags() {
			tags[t] = append(tags[t], name)
		}
	}

	tagList := []string{}

	for k, providers := range tags {
		if verbose {
			sort.Strings(providers)
			tagList = append(tagList, fmt.Sprintf("%s (%s)", k, strings.Join(providers, ", ")))
		} else {
			tagList = append(tagList, k)
		}
	}

	sort.Strings(tagList)

	return tagList
}

// Region returns the users region code.
// Eg. "US", "GB", etc
func Region() string {
	l := locale()

	tag, err := language.Parse(l)
	if err != nil {
		return defaultRegion
	}

	region, _ := tag.Region()

	return region.String()
}

// Language returns the users language code.
// Eg. "en", "es", etc
func Language() string {
	l := locale()

	tag, err := language.Parse(l)
	if err != nil {
		return defaultLanguage
	}

	base, _ := tag.Base()

	return base.String()
}

// SetClientLocale sets the locale of the client
// connecting to the web server.
func SetClientLocale(locale string) {
	clientLocale = locale
}

// locale gets the locale from the LANG env var if not set.
func locale() string {
	if clientLocale != "" {
		return clientLocale
	}

	lang := os.Getenv("LANG")
	if lang == "" {
		return ""
	}

	locale, _, _ := strings.Cut(lang, ".")

	return locale
}
