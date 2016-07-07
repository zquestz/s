package server

import (
	"bytes"
	"fmt"
	"net/http"
)

var (
	// indexTemplate is the go template for the index page.
	indexTemplate = `<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="apple-touch-icon" sizes="57x57" href="/apple-touch-icon-57x57.png">
    <link rel="apple-touch-icon" sizes="60x60" href="/apple-touch-icon-60x60.png">
    <link rel="apple-touch-icon" sizes="72x72" href="/apple-touch-icon-72x72.png">
    <link rel="apple-touch-icon" sizes="76x76" href="/apple-touch-icon-76x76.png">
    <link rel="apple-touch-icon" sizes="114x114" href="/apple-touch-icon-114x114.png">
    <link rel="apple-touch-icon" sizes="120x120" href="/apple-touch-icon-120x120.png">
    <link rel="apple-touch-icon" sizes="144x144" href="/apple-touch-icon-144x144.png">
    <link rel="apple-touch-icon" sizes="152x152" href="/apple-touch-icon-152x152.png">
    <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon-180x180.png">
    <link rel="icon" type="image/png" href="/favicon-32x32.png" sizes="32x32">
    <link rel="icon" type="image/png" href="/android-chrome-192x192.png" sizes="192x192">
    <link rel="icon" type="image/png" href="/favicon-96x96.png" sizes="96x96">
    <link rel="icon" type="image/png" href="/favicon-16x16.png" sizes="16x16">
    <link rel="manifest" href="/manifest.json">
    <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#272f32">
    <meta name="apple-mobile-web-app-title" content="s">
    <meta name="application-name" content="s">
    <meta name="msapplication-TileColor" content="#272f32">
    <meta name="msapplication-TileImage" content="/mstile-144x144.png">
    <meta name="theme-color" content="#272f32">
    <style>
{{.CSS}}
    </style>
    <title>s</title>
  </head>
  <body>
    <form name="search" action="/search" method="get">
      <div id="rightCorner">
        <select id="provider" name="provider" tabindex="2">
{{range .Providers}}           <option{{.|Label}}{{.|Selected}}>{{.}}</option>
{{end}}        </select><br>
        <select id="tag" name="tag" tabindex="3">
{{range .Tags}}           <option{{.|Label}}>{{.}}</option>
{{end}}        </select>
      </div>
      <input class="input" type="text" name="q" tabindex="1" placeholder="{{.Placeholder}}" autofocus required><br>
      <input type="submit" value="[ s ]" tabindex="4">
    </form>
    <script>
{{.JS}}
    </script>
  </body>
</html>
`

	// indexCSS is the css for the main index page.
	indexCSS = `*{font-family:"Tahoma","Geneva",sans-serif;font-size:14pt;text-align:center}
body{margin:0;padding:2em;background-color:#272F32;color:#DAEAEF}
#rightCorner{position:absolute;top:1.5em;right:1.5em;text-align:right}
a,a:visited{color:#FFF;font-size:.8em}
a:active,a:hover{color:#9DBDC6}
select{text-align:left;margin-bottom:.5em}
option{text-align:left}
form{margin-top:10em}
input[type=text]{width:100%;max-width:450px;border-bottom:1px solid #DAEAEF}
input{color:#DAEAEF;background-color:#272F32;display:inline-block;padding:0 0 .5em;margin:0 0 .5em;border:0;outline:none;border-radius:0;-webkit-border-radius:0;-webkit-appearance:none;appearance:none;-moz-appearance:none}
input:required{box-shadow:none}
input:invalid{box-shadow:none}
input[type=submit]{background-color:#272F32;font-size:2em;transition:color .2s ease}
input[type=submit]:active,input[type=submit]:hover{color:#9DBDC6}
input::-webkit-input-placeholder{font-style:italic}
input::-moz-placeholder{font-style:italic}
input:-moz-placeholder{font-style:italic}
input:-ms-input-placeholder{font-style:italic}
input:focus::-webkit-input-placeholder{color:transparent}
input:focus::-moz-placeholder{color:transparent}
input:focus:-moz-placeholder{color:transparent}
input:focus:-ms-input-placeholder{color:transparent}`
)

// indexJS is the javascript for the main index page.
func indexJS(defaultProvider string) string {
	var b bytes.Buffer

	b.WriteString(`function sInit() {
    "use strict";
    var searchForm = document.forms.search;
    searchForm.onsubmit = function () {
        var v = searchForm.q.value;
        if (v === null || v === "") {
            return false;
        }
    };
    var tag = document.getElementById("tag");
    var provider = document.getElementById("provider");
    tag.onchange = function () {
        if (tag.value === "") {
            if (provider.value === "") {
                provider.value = "`)

	b.WriteString(defaultProvider)
	b.WriteString(`";
            }
        } else {
            provider.value = "";
        }
    };
}
sInit();`)

	return b.String()
}

// searchTemplate is the html used for tag based searches.
func searchTemplate(uris []string, w http.ResponseWriter) {
	var b bytes.Buffer

	b.WriteString(`<!doctype html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>s</title>
  </head>
  <body>
    <script>
`)
	searchJS(uris, &b)
	b.WriteString(`
    </script>
  </body>
</html>
`)

	w.Header().Set("Content-Type", "text/html")
	w.Write(b.Bytes())
}

// SearchJS is the javascript used for tag based searches.
func searchJS(uris []string, b *bytes.Buffer) {
	b.WriteString(`function sInit() {
    "use strict";
    var urls = [];`)
	for _, u := range uris {
		b.WriteString(fmt.Sprintf("\n    urls.push(window.open(%q));", u))
	}
	b.WriteString(`
    var goBack = function () {
        window.history.back();
    };
    try {
        urls[urls.length - 1].focus();
        goBack();
    } catch (e) {
        alert("Pop-up Blocker is enabled! Please add this site to your exception list.");
        setTimeout(goBack, 5000);
    }
}
sInit();`)
}
