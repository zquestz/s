package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/zquestz/s/providers"
	"golang.org/x/text/language"
)

const (
	maxHeaderLength = 256
)

func search(defaultProvider string, verbose bool, w http.ResponseWriter, r *http.Request) {
	requestedProvider := r.FormValue("provider")
	requestedTag := r.FormValue("tag")
	locale := getLocale(r)

	if requestedProvider == "" && requestedTag == "" {
		requestedProvider = getProviderFromCookie(r, defaultProvider)
	}

	builders, err := buildProviderList(requestedProvider, requestedTag, locale)
	if err != nil {
		expandNotValid(err, w, r)
		return
	}

	if query := r.FormValue("q"); query != "" {
		handleSearch(builders, query, locale, verbose, w, r)
		return
	}

	queryNotFound(w, r)
}

func getLocale(r *http.Request) string {
	acceptLanguageHeader := r.Header.Get("Accept-Language")
	if strings.TrimSpace(acceptLanguageHeader) == "" {
		return providers.SystemLocale()
	}

	if len(acceptLanguageHeader) > maxHeaderLength {
		acceptLanguageHeader = acceptLanguageHeader[:maxHeaderLength]
	}

	if locale := parseAcceptLanguage(acceptLanguageHeader); locale != "" {
		return locale
	}

	return providers.SystemLocale()
}

func getProviderFromCookie(r *http.Request, defaultProvider string) string {
	if cookie, err := r.Cookie("s-provider"); err == nil && cookie.Value != "" {
		if _, exists := providers.Providers[cookie.Value]; exists {
			return cookie.Value
		}
	}
	return defaultProvider
}

func buildProviderList(requestedProvider, requestedTag, locale string) ([]providers.Provider, error) {
	var err error
	builders := []providers.Provider{}

	if requestedProvider != "" {
		requestedProvider, err = providers.ExpandProvider(requestedProvider)
		if err != nil {
			return nil, err
		}
		if provider := providers.Providers[requestedProvider]; provider != nil {
			builders = append(builders, provider)
		}
	}

	if requestedTag != "" {
		requestedTag, err = providers.ExpandTag(requestedTag, locale)
		if err != nil {
			return nil, err
		}
		builders = append(builders, providers.GetProvidersByTag(requestedTag, locale)...)
	}

	return builders, nil
}

func handleSearch(builders []providers.Provider, query, locale string, verbose bool, w http.ResponseWriter, r *http.Request) {
	uris := executeSearch(builders, query, locale)

	if verbose {
		for _, uri := range uris {
			fmt.Printf("%s\n", uri)
		}
	}

	if len(uris) == 1 {
		http.Redirect(w, r, uris[0], http.StatusFound)
		return
	}

	searchTemplate(uris, w)
}

func executeSearch(providerList []providers.Provider, query, locale string) []string {
	uris := make([]string, 0, len(providerList))

	for _, p := range providerList {
		uris = append(uris, p.BuildURI(query, locale))
	}

	return uris
}

func queryNotFound(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "A search query is required.", http.StatusBadRequest)
}

func expandNotValid(err error, w http.ResponseWriter, _ *http.Request) {
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
