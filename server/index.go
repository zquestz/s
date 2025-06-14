package server

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/zquestz/s/providers"
)

type templateVars struct {
	CSS         string
	JS          string
	Placeholder string
	Providers   []string
	Tags        []string
}

func index(defaultProvider string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	selectedProvider := defaultProvider
	if cookie, err := r.Cookie("s-provider"); err == nil && cookie.Value != "" {
		if _, exists := providers.Providers[cookie.Value]; exists {
			selectedProvider = cookie.Value
		}
	}

	t := template.New("index")

	selected := func(p string) string {
		if p == selectedProvider {
			return " selected"
		}

		return ""
	}

	label := func(opt string) string {
		if opt == "" {
			return fmt.Sprintf(" label=%q", "-")
		}

		return ""
	}

	t.Funcs(template.FuncMap{
		"Selected": selected,
		"Label":    label,
	})

	t, _ = t.Parse(indexTemplate)

	providerList := []string{""}
	providerList = append(providerList, providers.ProviderNames(false)...)

	tagList := []string{""}
	tagList = append(tagList, providers.TagNames(false)...)

	tvars := templateVars{
		CSS:         indexCSS,
		JS:          indexJS(selectedProvider),
		Placeholder: "kittens...",
		Providers:   providerList,
		Tags:        tagList,
	}

	t.Execute(w, tvars)
}
