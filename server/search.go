package server

import (
	"fmt"
	"net/http"

	"github.com/zquestz/s/providers"
)

func search(defaultProvider string, w http.ResponseWriter, r *http.Request) {
	requestedProvider := r.FormValue("provider")
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
