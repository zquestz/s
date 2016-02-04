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
		providerNotFound(w, r)
		return
	}

	if query := r.FormValue("q"); query != "" {
		uri := provider.BuildURI(query)
		http.Redirect(w, r, uri, 301)
		return
	}

	queryNotFound(w, r)
}

func providerNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	w.Write([]byte(fmt.Sprintf("Provider %q not found.", r.FormValue("provider"))))
}

func queryNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	w.Write([]byte("A search query is required."))
}
