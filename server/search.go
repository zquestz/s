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
	requestedTag := r.FormValue("tag")
	acceptLanguageHeader := r.Header.Get("Accept-Language")

	if strings.TrimSpace(acceptLanguageHeader) != "" {
		locale = parseAcceptLanguage(acceptLanguageHeader)
	}

	if requestedProvider == "" && requestedTag == "" {
		requestedProvider = defaultProvider
	}

	provider := providers.Providers[requestedProvider]
	if provider == nil && requestedTag == "" {
		providerNotFound(requestedProvider, w, r)
		return
	}

	builders := []providers.Provider{}

	if provider != nil {
		builders = append(builders, provider)
	}

	if requestedTag != "" {
		builders = append(builders, providers.GetProvidersByTag(requestedTag)...)
	}

	if len(builders) == 0 {
		tagProvidersNotFound(requestedTag, w, r)
		return
	}

	if query := r.FormValue("q"); query != "" {
		uris := executeSearch(builders, query, locale)

		if verbose {
			for _, uri := range uris {
				fmt.Printf("%s\n", uri)
			}
		}

		if len(uris) == 1 {
			http.Redirect(w, r, uris[0], 301)
			return
		}

		searchTemplate(uris, w)
		return
	}

	queryNotFound(w, r)
}

func executeSearch(providerList []providers.Provider, query, locale string) []string {
	uris := make([]string, 0, len(providerList))

	clientLocaleMutex.Lock()
	defer clientLocaleMutex.Unlock()

	providers.SetClientLocale(locale)

	for _, p := range providerList {
		uris = append(uris, p.BuildURI(query))
	}

	return uris
}

func providerNotFound(provider string, w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("Provider %q not found.", provider), 400)
}

func tagProvidersNotFound(tag string, w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("No providers matched tag %q.", tag), 400)
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
