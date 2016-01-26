package providers

import (
	"fmt"
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
	builder := Providers[p]

	if builder != nil {
		url := builder.BuildURI(q)
		if verbose {
			fmt.Printf("%s\n", url)
		}
		launcher.OpenURI(binary, url)
	} else {
		fmt.Printf("Provider %q not supported!\n", p)
	}
}

// DisplayProviders displays all the loaded providers.
func DisplayProviders() string {
	names := []string{}

	for key, _ := range Providers {
		names = append(names, key)
	}

	sort.Strings(names)

	return fmt.Sprintf("%s\n", strings.Join(names, "\n"))
}
