package gopkg

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("godoc", &GoDocProvider{})
}

// GoDocProvider adheres to the Provider interface.
type GoDocProvider struct {
}

// BuildURI generates a search URL for Go.
func (p *GoDocProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://godoc.org/?q=%s", url.QueryEscape(q))
}
