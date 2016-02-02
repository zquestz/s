package server

import (
	"net/http"
	"text/template"

	"github.com/zquestz/s/providers"
)

var (
	tmpl = `<!doctype html>
<html>
  <head>
    <link rel="stylesheet" href="http://yui.yahooapis.com/pure/0.6.0/pure-min.css">
    <style>
    body {
      padding: 3em;
      font-size: 16px;
    }

    h1, form, div {
      text-align: center;
    }

    h1 {
      color: white;
    }

    body {
        background-color: black;
    }

    input, select {
      padding: 8px;
    }

    .footer {
      padding: 20px;
      font-size: 1em;
    }
    </style>
    <title>s</title>
  </head>
  <body>
    <h1>s</h1>
    <form action="/search" method="POST">
      <input type="text" name="search" tabindex="1" autofocus />
      <select name="provider">
      {{range .Providers}}<option{{.|Selected}}>{{.}}</option>{{end}}
      </select>
    </form>
    <div class="footer"><a href="https://github.com/zquestz/s">https://github.com/zquestz/s</a></div>
  </body>
</html>
`
)

type templateVars struct {
	Providers       []string
	defaultProvider string
}

func index(provider string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	t := template.New("search")

	selected := func(p string) string {
		if p == provider {
			return " selected"
		}

		return ""
	}

	t.Funcs(template.FuncMap{
		"Selected": selected,
	})

	t, _ = t.Parse(tmpl)

	tvars := templateVars{
		Providers:       providers.ProviderNames(),
		defaultProvider: provider,
	}

	t.Execute(w, tvars)
}
