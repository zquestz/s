package mdn

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("mdn", &MDNProvider{})
}

//MDNProvider adheres to the Provider interface.
type MDNProvider struct {
}

//BuildURI generates a search URL for MDN.
func (p *MDNProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://developer.mozilla.org/en-US/search?q=%s", url.QueryEscape(q))
}
