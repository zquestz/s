package server

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/zquestz/s/providers"
	"golang.org/x/text/language"
)

var (
	clientLocaleMutex sync.Mutex
)

func search(defaultProvider string, verbose bool, w http.ResponseWriter, r *http.Request) {
	var locale string

	requestedProvider := r.FormValue("provider")
	acceptLanguageHeader := r.Header.Get("Accept-Language")

	if strings.TrimSpace(acceptLanguageHeader) != "" {
		locale = parseAcceptLanguage(acceptLanguageHeader)
	}

	if requestedProvider == "" {
		requestedProvider = defaultProvider
	}

	provider := providers.Providers[requestedProvider]
	if provider == nil {
		providerNotFound(requestedProvider, w, r)
		return
	}

	if query := r.FormValue("q"); query != "" {
		uri := executeSearch(provider, query, locale)

		if verbose {
			fmt.Printf("%s\n", uri)
		}

		http.Redirect(w, r, uri, 301)
		return
	}

	queryNotFound(w, r)
}

func executeSearch(provider providers.Provider, query, locale string) string {
	clientLocaleMutex.Lock()
	defer clientLocaleMutex.Unlock()

	providers.SetClientLocale(locale)

	return provider.BuildURI(query)
}

func providerNotFound(provider string, w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("Provider %q not found.", provider), 400)
}

func queryNotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("A search query is required."), 400)
}

func parseAcceptLanguage(header string) string {
	tags, _, err := language.ParseAcceptLanguage(header)
	if err != nil {
		return ""
	}

	// Using the first Tag as Tags are sorted by highest weight first
	return tags[0].String()
}
