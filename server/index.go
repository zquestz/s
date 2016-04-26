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

	t := template.New("index")

	selected := func(p string) string {
		if p == defaultProvider {
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
		JS:          indexJS(defaultProvider),
		Placeholder: "kittens...",
		Providers:   providerList,
		Tags:        tagList,
	}

	t.Execute(w, tvars)
}
