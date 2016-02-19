[![License][License-Image]][License-Url] [![ReportCard][ReportCard-Image]][ReportCard-Url] [![Build][Build-Status-Image]][Build-Status-Url] [![Release][Release-Image]][Release-Url]
# s
Web search from the terminal. Just opens in your browser.

```
Usage:
  s <query> [flags]

Flags:
  -b, --binary string     binary to launch search URI
  -c, --cert string       path to cert.pem for TLS
  -k, --key string        path to key.pem for TLS
  -l, --list-providers    list supported providers
      --port int          server port (default 8080)
  -p, --provider string   search provider (default "google")
  -s, --server            launch web server
  -v, --verbose           display URL when opening
      --version           display version
```

## Install

```
go get -v github.com/zquestz/s
cd $GOPATH/src/github.com/zquestz/s
make
make install
```

## Examples

Search for puppies on google.
```
s puppies
```

Search for a wifi router on amazon
```
s -p amazon wifi router
```

Search for rhinos on wikipedia
```
s -p wikipedia rhinos
```

## Provider Expansion

We can do partial matching of provider names. This searches Facebook for hamsters.
```
s -p fa hamsters
```

Or toasters on amazon.
```
s -p am toasters
```

## Provider Autocompletion

Autocompletion for providers is supported. For setting up autocompletion:

1. Have `s` installed
2. Add the following lines to `~/.bash_profile` or `~/.zshrc`
```
if [ -f $GOPATH/src/github.com/zquestz/s/autocomplete/s-completion.bash ]; then
    . $GOPATH/src/github.com/zquestz/s/autocomplete/s-completion.bash
fi
```

Now you are good to go.
```
s -p ba<TAB><TAB>
baidu     bandcamp
```

### Fish

Alternatively, if you use [fish](http://fishshell.com/), the following will work:
```
mkdir -p ~/.config/fish/completions
ln -s $GOPATH/src/github.com/zquestz/s/autocomplete/s.fish ~/.config/fish/completions/s.fish
```

## Advanced

Setup an alias in your `.profile` for your favorite providers.
```
alias sa="s -p amazon"
alias sw="s -p wikipedia"
```

Use w3m to find cats instead of just your default browser.
```
s -b w3m cats
```

Search for conspiracy theories in incognito mode.
```
s -b "chromium --incognito" conspiracy theories
s -b "firefox --private-window" conspiracy theories
```

Search in a specific subreddit.
```
s -p reddit /r/cscareerquestions best startups.
```

## Server Mode

A web interface is also provided. Just pass the `-s` flag.

Start a server on port 8080 (default).
```
s -s
```

Start a server with TLS on port 8443.
```
s -s -c /path/to/cert.pem -k /path/to/key.pem --port 8443
```

Feel free to try it out at [https://jumps.io/](https://jumps.io/).

## Configuration

To setup your own default configuration just create `~/.s/config`. The configuration file is in UCL format. JSON is also fully supported as UCL can parse JSON files.

For more information about UCL visit:
[https://github.com/vstakhov/libucl](https://github.com/vstakhov/libucl)

The following keys are supported:

* blacklist (array of providers to exclude)
* binary (binary to launch search URI)
* cert (path to cert.pem for TLS)
* customProviders (array of custom providers)
* key (path to key.pem for TLS)
* port (server port)
* provider (search provider)
* verbose (display URL when opening)
* whitelist (array of providers to include)

Set your default provider to duckduckgo:
```
provider: duckduckgo
```

To only search a few providers:
```
whitelist: [google, amazon, wikipedia]
```

To exclude providers you don't need:
```
blacklist: [dumpert]
```

To add a custom provider:
```
customProviders [
  {
    name: example
    url: "http://example.com?q=%s"
  }
]
```

Custom providers require a few things:
* An alphanumeric name. `^[a-zA-Z0-9_]*$`
* A `%s` token for the query string.
* A valid url scheme.

## Supported Providers

* 500px
* 8tracks
* amazon
* arstechnica
* arxiv
* atmospherejs
* baidu
* bandcamp
* bing
* codepen
* coursera
* cplusplus
* crunchyroll
* debianpkg
* digg
* diigo
* dockerhub
* dribbble
* duckduckgo
* dumpert
* facebook
* flickr
* flipkart
* foursquare
* gist
* github
* gmail
* go
* godoc
* google
* googleplus
* hackernews
* ietf
* ifttt
* imdb
* imgur
* inbox
* instagram
* kickasstorrents
* libgen
* linkedin
* lmgtfy
* macports
* mdn
* msdn
* naver
* netflix
* nhaccuatui
* npm
* npmsearch
* nvd
* overstock
* packagist
* php
* pinterest
* python
* quora
* reddit
* rottentomatoes
* rubygems
* soundcloud
* spotify
* stackoverflow
* steam
* taobao
* thepiratebay
* torrentz
* twitchtv
* twitter
* unity3d
* upcloud
* vimeo
* wikipedia
* wolframalpha
* yahoo
* yandex
* youtube

#### Contributors

* [Josh Ellithorpe (zquestz)](https://github.com/zquestz/)
* [Christian Petersen (fnky)](https://github.com/fnky/)
* [Preet Bhinder (mbhinder)](https://github.com/mbhinder/)
* [Robert-Jan Keizer (KeizerDev)](https://github.com/KeizerDev/)
* [Vitor Cortez (vekat)](https://github.com/vekat/)
* [David Liu (tw4dl)](https://github.com/tw4dl/)
* [Lex Broner (akb)](http://github.com/akb/)
* [Diego Jara (djap96)](https://github.com/djap96/)
* [Luvsandondov Lkhamsuren (lkhamsurenl)](https://github.com/lkhamsurenl/)
* [Eray Aydın (erayaydin)](https://github.com/erayaydin/)
* [Murilo Santana (mvrilo)](https://github.com/mvrilo/)

#### License

s is released under the MIT license.

[License-Url]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
[ReportCard-Url]: http://goreportcard.com/report/zquestz/s
[ReportCard-Image]: http://goreportcard.com/badge/zquestz/s
[Build-Status-Url]: http://travis-ci.org/zquestz/s
[Build-Status-Image]: https://travis-ci.org/zquestz/s.svg?branch=master
[Release-Url]: https://github.com/zquestz/s/releases/tag/v0.4.2
[Release-image]: http://img.shields.io/badge/release-v0.4.2-1eb0fc.svg
