package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/zquestz/s/providers"
	"golang.org/x/text/language"
)

func search(defaultProvider string, verbose bool, w http.ResponseWriter, r *http.Request) {
	requestedProvider := r.FormValue("provider")
	acceptLanguageHeader := r.Header.Get("Accept-Language")

	if strings.TrimSpace(acceptLanguageHeader) != "" {
		locale := parseAcceptLanguage(acceptLanguageHeader)
		providers.SetClientLocale(locale)
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
		uri := provider.BuildURI(query)

		if verbose {
			fmt.Printf("%s\n", uri)
		}

		http.Redirect(w, r, uri, 301)
		return
	}

	queryNotFound(w, r)
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
