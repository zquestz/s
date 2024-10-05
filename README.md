# s

[![License][License-Image]][License-URL] [![ReportCard][ReportCard-Image]][ReportCard-URL] [![Build][Build-Status-Image]][Build-Status-URL] [![Release][Release-Image]][Release-URL] [![🧪 Snap Builds](https://github.com/kz6fittycent/s/actions/workflows/test-snap-can-build.yml/badge.svg)](https://github.com/kz6fittycent/s/actions/workflows/test-snap-can-build.yml) <a href="https://snapcraft.io/s-search"><img src="https://snapcraft.io/s-search/badge.svg" alt="Snap Status"></a>

Web search from the terminal. Just opens in your browser.

```text
Usage:
  s <query> [flags]

Flags:
  -b, --binary string       binary to launch search URI
  -c, --cert string         path to cert.pem for TLS
      --completion string   completion script for bash, zsh, fish or powershell
  -h, --help                help for s
  -k, --key string          path to key.pem for TLS
  -l, --list-providers      list supported providers
      --list-tags           list available tags
  -o, --output              output only mode
      --port int            server port (default 8080)
  -p, --provider string     search provider (default "presearch")
  -s, --server              launch web server
  -t, --tag string          search tag
  -v, --verbose             verbose mode
      --version             display version
```

## Install

Make sure that `GOPATH` and `GOBIN` env vars are set. Then run:

```zsh
go install github.com/zquestz/s@latest
```

Alternatively, you can use Homebrew:

```zsh
brew install s-search
```

Arch Linux users can install from the AUR:

```zsh
yay -S s
```

## Examples

Search for puppies on presearch.

```zsh
s puppies
```

Search for dragonflies on google.

```zsh
s -p google dragonflies
```

Search for a wifi router on amazon

```zsh
s -p amazon wifi router
```

Search for rhinos on wikipedia

```zsh
s -p wikipedia rhinos
```

Search providers tagged "video" for muppets.

```zsh
s -t video muppets
```

## Provider/Tag Expansion

We can do partial matching of provider and tag names. This searches Facebook for hamsters.

```zsh
s -p fa hamsters
```

Or toasters on amazon.

```zsh
s -p am toasters
```

This searches "tech-news" tagged providers for ssd info.

```zsh
s -t te ssd
```

Or shopping sites for blankets.

```zsh
s -t sh blankets
```

## Provider/Tag Autocompletion

Autocompletion is supported for providers and tags. To set up autocompletion:

### Bash Linux

```zsh
s --completion bash > /etc/bash_completion.d/s
```

### Bash MacOS

```zsh
s --completion bash > /usr/local/etc/bash_completion.d/s
```

### Zsh

Generate a `_s` completion script and put it somewhere in your `$fpath`:

```zsh
s --completion zsh > /usr/local/share/zsh/site-functions/_s
```

### Fish

```zsh
s --completion fish > ~/.config/fish/completions/s.fish
```

### Powershell

```powershell
(& s --completion powershell) | Out-String | Invoke-Expression
```

## Advanced

Setup an alias in your `.profile` for your favorite providers.

```bash
alias sa="s -p amazon"
alias sw="s -p wikipedia"
```

Use w3m to find cats instead of just your default browser.

```zsh
s -b w3m cats
```

Search for conspiracy theories in incognito mode.

```zsh
s -b "chromium --incognito" conspiracy theories
s -b "firefox --private-window" conspiracy theories
```

Search in a specific subreddit.

```zsh
s -p reddit /r/cscareerquestions best startups.
```

## Server Mode

A web interface is also provided. Just pass the `-s` flag.

Start a server on port 8080 (default).

```zsh
s -s
```

Start a server with TLS on port 8443.

```zsh
s -s -c /path/to/cert.pem -k /path/to/key.pem --port 8443
```

Feel free to try it out at [https://jumps.io/](https://jumps.io/).

## Configuration

To setup your own default configuration just create `~/.config/s/config`. The configuration file is in UCL format.

For more information about UCL visit:
[https://github.com/vstakhov/libucl](https://github.com/vstakhov/libucl)

The following keys are supported:

* blacklist (array of providers to exclude)
* binary (binary to launch search URI)
* cert (path to cert.pem for TLS)
* customProviders (array of custom providers)
* key (path to key.pem for TLS)
* output (output only mode)
* port (server port)
* provider (search provider)
* tag (search tag)
* verbose (verbose mode)
* whitelist (array of providers to include)

Set your default provider to duckduckgo:

```ucl
provider: duckduckgo
```

To only search a few providers:

```ucl
whitelist: [google, amazon, wikipedia]
```

To exclude providers you don't need:

```ucl
blacklist: [dumpert]
```

To add a custom provider:

```ucl
customProviders [
  {
    name: example
    url: "http://example.com?q=%s"
    tags: [example]
  }
]
```

Custom providers require a few things:

* An alphanumeric name. `^[a-zA-Z0-9_]*$`
* A `%s` token for the query string.
* A valid URL scheme.

## Supported Providers

* 500px
* 8tracks
* aliexpress
* allocine
* amazon
* archpkg
* archwiki
* ardmediathek
* arstechnica
* arxiv
* atmospherejs
* aur
* baidu
* bandcamp
* bgr
* bigbasket
* bing
* brave
* buzzfeed
* cnn
* codepen
* coursera
* cplusplus
* cppreference
* crates
* crunchyroll
* debianpkg
* dict
* diigo
* dockerhub
* dribbble
* duckduckgo
* dumpert
* ecosia
* engadget
* explainshell
* facebook
* flake8rules
* flickr
* flipkart
* foursquare
* freebsdman
* freshports
* gibiru
* giphy
* gist
* github
* gmail
* go
* godoc
* goodreads
* google
* googledocs
* hackernews
* idealo
* ietf
* ifttt
* imdb
* imgur
* instagram
* kagi
* kaufda
* libgen
* linkedin
* lmgtfy
* macports
* magnetdl
* mdn
* medium
* metacpan
* msdn
* naver
* netflix
* nhaccuatui
* npm
* npr
* nvd
* openbsdman
* overstock
* packagist
* perplexity
* presearch
* protondb
* phandroid
* phind
* php
* pinterest
* postgresql
* pydoc
* pypi
* python
* quora
* qwant
* reddit
* regex
* rottentomatoes
* rubygems
* shodan
* soundcloud
* spotify
* stackoverflow
* steam
* taobao
* thepiratebay
* theregister
* twitchtv
* twitter
* ultimateguitar
* unity3d
* vimeo
* wikipedia
* wolframalpha
* yahoo
* yandex
* youtube
* zdf
* zhihu

## Contributors

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
* [Jun He (knarfeh)](https://github.com/knarfeh/)
* [Xavier Bruhiere (hackliff)](github.com/hackliff)

## License

s is released under the MIT license.

[License-URL]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
[ReportCard-URL]: http://goreportcard.com/report/zquestz/s
[ReportCard-Image]: https://goreportcard.com/badge/github.com/zquestz/s
[Build-Status-URL]: https://app.travis-ci.com/github/zquestz/s
[Build-Status-Image]: https://app.travis-ci.com/zquestz/s.svg?branch=master
[Release-URL]: https://github.com/zquestz/s/releases/tag/v0.7.1
[Release-Image]: http://img.shields.io/badge/release-v0.7.1-1eb0fc.svg
