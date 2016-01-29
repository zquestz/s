package imdb

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("imdb", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for IMDB.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.imdb.com/find?q=%s", url.QueryEscape(q))
}
