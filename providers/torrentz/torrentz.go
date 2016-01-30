package torrentz

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("torrentz", &TorrentzProvider{})
}

type TorrentzProvider struct {
}

// BuildURI generates a search URL for Torrentz.
func (p *TorrentzProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.torrentz.eu/search?q=%s", url.QueryEscape(q))
}
