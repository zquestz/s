package netflix

import (
	"fmt"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("netflix", &NetflixProvider{})
}

// NetflixProvider adheres to the Provider interface.
type NetflixProvider struct {
}

// BuildURI generates a search URL for Netflix.
func (p *NetflixProvider) BuildURI(q string) string {
	return fmt.Sprintf("http://www.netflix.com/search/%s", q)
}
