package gopkg

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("go", &GoProvider{})
}

// GoProvider adheres to the Provider interface.
type GoProvider struct {
}

// BuildURI generates a search URL for Go.
func (p *GoProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://golang.org/search?q=%s", url.QueryEscape(q))
}
