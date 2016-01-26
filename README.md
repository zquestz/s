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

## Supported Providers

* amazon
* bing
* digg
* dockerhub
* duckduckgo
* facebook
* github
* gist
* go
* godoc
* google
* npm
* npmsearch
* pinterest
* quora
* reddit
* soundcloud
* stackoverflow
* twitter
* wikipedia
* yahoo
* youtube

## Install

```
go get -v github.com/zquestz/s
cd $GOPATH/src/github.com/zquestz/s
make
make install
```

## Provider Expansion

We can do partial matching of provider names. This searches Facebook for hamsters.
```
s -p f hamsters
```

Or toasters on amazon.
```
s -p a toasters
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

#### Contributors

* [Josh Ellithorpe (zquestz)](https://github.com/zquestz/)
* [Christian Petersen (fnky)](https://github.com/fnky/)
* [Preet Bhinder (mbhinder)](https://github.com/mbhinder/)
* [Vitor Cortez (vekat)](https://github.com/vekat/)

#### License

s is released under the MIT license.
