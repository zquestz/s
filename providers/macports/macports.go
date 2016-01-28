package macports

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("macports", &MacPortsProvider{})
}

// MacPortsProvider adheres to the Provider interface.
type MacPortsProvider struct {
}

// BuildURI generates a search URL for macports.
func (p *MacPortsProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.macports.org/ports.php?by=name&substr=%s", url.QueryEscape(q))
}
