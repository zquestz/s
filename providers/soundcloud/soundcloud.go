package soundcloud

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("soundcloud", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for SoundCloud.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://soundcloud.com/search?q=%s", url.QueryEscape(q))
}
