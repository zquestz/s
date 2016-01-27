package atmospherejs

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("atmospherejs", &AtmosphereJSProvider{})
}

// AtmosphereJSProvider adheres to the Provider interface.
type AtmosphereJSProvider struct {
}

// BuildURI generates a search URL for AtmosphereJS.
func (p *AtmosphereJSProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://atmospherejs.com/?q=%s", url.QueryEscape(q))
}
