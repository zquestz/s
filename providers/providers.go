package providers

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/zquestz/s/launcher"
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

// Search builds a search URL and opens it in your browser.
func Search(binary string, p string, q string, verbose bool) {
	prov, err := ExpandProvider(p)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
		os.Exit(1)
	}

	builder := Providers[prov]

	if builder != nil {
		url := builder.BuildURI(q)
		if verbose {
			fmt.Printf("%s\n", url)
		}
		launcher.OpenURI(binary, url)
	} else {
		fmt.Printf("Provider %q not supported!\n", prov)
	}
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
		if r.Match([]byte(n)) {
			validProviders = append(validProviders, n)
		}
	}

	switch len(validProviders) {
	case 0:
		return "", fmt.Errorf("No valid provider found for %q", provider)
	case 1:
		return validProviders[0], nil
	default:
		return "", fmt.Errorf("Multiple providers matched %q: %v", provider, validProviders)
	}
}

// ProviderNames returns a sorted slice of provider names.
func ProviderNames() []string {
	names := []string{}

	for key, _ := range Providers {
		names = append(names, key)
	}

	sort.Strings(names)
	return names
}
