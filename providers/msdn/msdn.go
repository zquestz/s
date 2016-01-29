package msdn

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("msdn", &Provider{})
}

type Provider struct {}

// BuildURI generates a search URL for MSDN.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://social.msdn.microsoft.com/Search/en-US?query=%s", url.QueryEscape(q))
}
