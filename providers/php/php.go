package php

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("php", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for Php.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://php.net/%s", url.QueryEscape(q))
}
