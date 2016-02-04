package server

var (
	// IndexTemplate is the go template for the index page.
	IndexTemplate = `<!doctype html>
<html>
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <style>
*{font-family:"Tahoma","Geneva",sans-serif;font-size:14pt;text-align:center}body{margin:0;padding:2em;background-color:#272F32;color:#DAEAEF}a,a:visited{color:#FFF;font-size:.8em}a:active,a:hover{color:#9DBDC6}select{position:absolute;top:1.5em;right:1.5em;text-align:left}option{text-align:left}form{margin-top:10em}input[type=text]{width:100%;max-width:450px;border-bottom:1px solid #DAEAEF}input{color:#DAEAEF;background-color:#272F32;display:inline-block;padding-bottom:.5em;margin-bottom:.5em;border:0;outline:none}input:required{box-shadow:none}input:invalid{box-shadow:none}input[type=submit]{background-color:#272F32;font-size:2em;transition:color .2s ease;letter-spacing:.2em;padding-left:.2em}input[type=submit]:active,input[type=submit]:hover{color:#9DBDC6}input::-webkit-input-placeholder{font-style:italic}input::-moz-placeholder{font-style:italic}input:-moz-placeholder{font-style:italic}input:focus:-ms-input-placeholder{font-style:italic}input:focus::-webkit-input-placeholder{color:transparent}input:focus::-moz-placeholder{color:transparent}input:focus:-moz-placeholder{color:transparent}input:focus:-ms-input-placeholder{color:transparent}
    </style>
    <title>s</title>
  </head>
  <body>
    <form action="/search" method="POST">
      <select name="provider" tabindex="2">
{{range .Providers}}         <option{{.|Selected}}>{{.}}</option>
{{end}}      </select><br />
      <input class="input" type="text" name="q" tabindex="1" placeholder="{{.Placeholder}}" autofocus required /><br />
      <input type="submit" value="[s]" tabindex="3" />
    </form>
  </body>
</html>
`
)
