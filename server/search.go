package server

import (
	"fmt"
	"net/http"

	"github.com/zquestz/s/providers"
)

func search(w http.ResponseWriter, r *http.Request) {
	provider := providers.Providers[r.FormValue("provider")]
	if provider == nil {
		notFound(w, r)
		return
	}

	uri := provider.BuildURI(r.FormValue("q"))

	http.Redirect(w, r, uri, 301)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	w.Write([]byte(fmt.Sprintf("Provider %q not found.", r.FormValue("provider"))))
}
