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
	var err error

	requestedProvider := r.FormValue("provider")
	requestedTag := r.FormValue("tag")
	acceptLanguageHeader := r.Header.Get("Accept-Language")

	if strings.TrimSpace(acceptLanguageHeader) != "" {
		locale = parseAcceptLanguage(acceptLanguageHeader)
	}

	if requestedProvider == "" && requestedTag == "" {
		requestedProvider = defaultProvider
	}

	if requestedProvider != "" {
		requestedProvider, err = providers.ExpandProvider(requestedProvider)
		if err != nil {
			expandNotValid(err, w, r)
			return
		}
	}

	provider := providers.Providers[requestedProvider]
	builders := []providers.Provider{}

	if provider != nil {
		builders = append(builders, provider)
	}

	if requestedTag != "" {
		requestedTag, err = providers.ExpandTag(requestedTag)
		if err != nil {
			expandNotValid(err, w, r)
			return
		}

		builders = append(builders, providers.GetProvidersByTag(requestedTag)...)
	}

	if query := r.FormValue("q"); query != "" {
		uris := executeSearch(builders, query, locale)

		if verbose {
			for _, uri := range uris {
				fmt.Printf("%s\n", uri)
			}
		}

		if len(uris) == 1 {
			http.Redirect(w, r, uris[0], http.StatusMovedPermanently)
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

func queryNotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("A search query is required."), http.StatusBadRequest)
}

func expandNotValid(err error, w http.ResponseWriter, r *http.Request) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func parseAcceptLanguage(header string) string {
	tags, _, err := language.ParseAcceptLanguage(header)
	if err != nil {
		return ""
	}

	// Using the first Tag as Tags are sorted by highest weight first
	return tags[0].String()
}
