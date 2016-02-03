package server

import (
	"net/http"

	"github.com/zquestz/s/providers"
)

func search(w http.ResponseWriter, r *http.Request) {
	provider := providers.Providers[r.FormValue("provider")]
	uri := provider.BuildURI(r.FormValue("q"))

	http.Redirect(w, r, uri, 301)
}
