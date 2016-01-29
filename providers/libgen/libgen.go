package libgen

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("libgen", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for Libgen.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("http://gen.lib.rus.ec/search.php?req=%s", url.QueryEscape(q))
}
