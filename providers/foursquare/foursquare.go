package foursquare

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("foursquare", &Provider{})
}

type Provider struct{}

//BuildURI generates a search URL for Foursquare.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://foursquare.com/explore?&q=%s", url.QueryEscape(q))
}
