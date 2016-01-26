package gopkg

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("php", &PhpProvider{})
}

// PhpProvider adheres to the Provider interface.
type PhpProvider struct {
}

// BuildURI generates a search URL for Php.
func (p *PhpProvider) BuildURI(q string) string {
	return fmt.Sprintf("http://php.net/%s", url.QueryEscape(q))
}
