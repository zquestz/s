package nvd

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("nvd", &NVDProvider{})
}

// NVDProvider adheres to the Provider interface.
type NVDProvider struct {
}

// BuildURI generates a search URL for the National Vulnerability Database.
func (p *NVDProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://web.nvd.nist.gov/view/vuln/search-results?query=%s", url.QueryEscape(q))
}
