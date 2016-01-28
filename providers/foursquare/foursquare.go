package foursquare

import (
	"fmt"
	"net/url"
	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("foursquare", &FoursquareProvider{})
}

//FoursquareProvider adheres to the Provider interface.
type FoursquareProvider struct {
}

//BuildURI generates a search URL for Foursquare.
func (p *FoursquareProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://foursquare.com/explore?&q=%s", url.QueryEscape(q))
}
