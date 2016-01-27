package libgen

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("libgen", &Libgen{})
}

// Libgen adheres to the Provider interface.
type Libgen struct {
}

// BuildURI generates a search URL for Libgen.
func (p *Libgen) BuildURI(q string) string {
	return fmt.Sprintf("http://gen.lib.rus.ec/search.php?req=%s", url.QueryEscape(q))
}
