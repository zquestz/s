package imdb

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("imdb", &IMDBProvider{})
}

// IMDBProvider adheres to the Provider interface.
type IMDBProvider struct {
}

// BuildURI generates a search URL for IMDB.
func (p *IMDBProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.imdb.com/find?q=%s", url.QueryEscape(q))
}
