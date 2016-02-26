package metacpan

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("metacpan", &MetaCPANProvider{})
}

// MetaCPANProvider adheres to the Provider interface.
type MetaCPAN struct{
}

// BuildURI generates a search URL for MetaCPAN.
func (p *MetaCPANProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://metacpan.org/search?q=%s", url.QueryEscape(q))
}
