package ietf

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("ietf", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for IETF.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://datatracker.ietf.org/doc/search/"+
		"?name=%s&rfcs=on&activedrafts=on", url.QueryEscape(q))
}
