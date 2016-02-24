package providers

import (
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
	blacklist       = make(map[string]interface{})
	whitelist       = make(map[string]interface{})
)

// Provider interface provides a way to build the URI
// for each provider.
type Provider interface {
	BuildURI(string) string
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
func Search(binary string, p string, q string, verbose bool) error {
	prov, err := ExpandProvider(p)
	if err != nil {
		return err
	}

	builder := Providers[prov]

	if builder != nil {
		url := builder.BuildURI(q)

		if verbose {
			fmt.Printf("%s\n", url)
		}

		return launcher.OpenURI(binary, url)
	}

	return fmt.Errorf("Provider %q not supported!\n", prov)
}

// DisplayProviders displays all the loaded providers.
func DisplayProviders() string {
	names := ProviderNames()

	return fmt.Sprintf("%s\n", strings.Join(names, "\n"))
}

// ExpandProvider expands the passed in provider to the full value.
func ExpandProvider(provider string) (string, error) {
	names := ProviderNames()
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
		return "", fmt.Errorf("No provider found for %q", provider)
	case 1:
		return validProviders[0], nil
	default:
		return "", fmt.Errorf("Multiple providers matched %q: %v", provider, validProviders)
	}
}

// ProviderNames returns a sorted slice of provider names.
func ProviderNames() []string {
	names := []string{}

	for key := range Providers {
		if enableWhitelist {
			_, ok := whitelist[key]
			if !ok {
				continue
			}
		}

		if enableBlacklist {
			_, ok := blacklist[key]
			if ok {
				continue
			}
		}

		names = append(names, key)
	}

	sort.Strings(names)
	return names
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

func locale() string {
	if clientLocale != "" {
		return clientLocale
	}

	lang := os.Getenv("LANG")
	if lang == "" {
		return ""
	}

	locale := strings.Split(lang, ".")[0]

	return locale
}
