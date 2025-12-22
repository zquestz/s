package overstock

import (
	"fmt"
	"strings"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("overstock", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Overstock.
func (p *Provider) BuildURI(q string) string {
	query := strings.ReplaceAll(q, " ", "-")
	return fmt.Sprintf("https://www.overstock.com/%s,/k,/results.html", query)
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"shopping"}
}
