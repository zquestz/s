package dumpert

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("dumpert", &DumpertProvider{})
}

// DumpertProvider adheres to the Provider interface.
type DumpertProvider struct {
}

// BuildURI generates a search URL for dumpert.
func (p *DumpertProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.dumpert.nl/search/ALL/%s/", url.QueryEscape(q))
}
