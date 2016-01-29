package eighttracks

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("8tracks", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for 8tracks.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://8tracks.com/explore/%s", url.QueryEscape(q))
}
