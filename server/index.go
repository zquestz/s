package server

import (
	"net/http"
	"text/template"

	"github.com/zquestz/s/providers"
)

type templateVars struct {
	Providers       []string
	defaultProvider string
}

func index(defaultProvider string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	t := template.New("index")

	selected := func(p string) string {
		if p == defaultProvider {
			return " selected"
		}

		return ""
	}

	t.Funcs(template.FuncMap{
		"Selected": selected,
	})

	t, _ = t.Parse(IndexTemplate)

	tvars := templateVars{
		Providers: providers.ProviderNames(),
	}

	t.Execute(w, tvars)
}
