[![License][License-Image]][License-Url] [![Build][Build-Status-Image]][Build-Status-Url] [![Release][Release-Image]][Release-Url]
# s
Web search from the terminal. Just opens your default browser.

```
Usage:
  s <query> [flags]

Flags:
  -b, --binary string     binary to launch search uri
  -l, --list-providers    list supported providers
  -p, --provider string   set search provider (default "google")
  -v, --verbose           display url when opening
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
```

## Supported Providers

* 8tracks
* amazon
* arxiv
* atmospherejs
* baidu
* bandcamp
* bing
* codepen
* coursera
* crunchyroll
* digg
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
* hackernews
* ietf
* ifttt
* imdb
* imgur
* kickasstorrents
* libgen
* linkedin
* macports
* mdn
* msdn
* nhaccuatui
* npm
* npmsearch
* nvd
* packagist
* php
* pinterest
* python
* quora
* reddit
* rubygems
* soundcloud
* spotify
* stackoverflow
* steam
* taobao
* thepiratebay
* twitchtv
* twitter
* unity3d
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

#### License

s is released under the MIT license.

[License-Url]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
[Build-Status-Url]: http://travis-ci.org/zquestz/s
[Build-Status-Image]: https://travis-ci.org/zquestz/s.svg?branch=master
[Release-Url]: https://github.com/zquestz/s/releases/tag/v0.2.0
[Release-image]: http://img.shields.io/badge/release-v0.2.0-1eb0fc.svg
