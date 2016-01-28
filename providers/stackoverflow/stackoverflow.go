package stackoverflow

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("stackoverflow", &StackOverflowProvider{})
}

// StackOverflowProvider adheres to the Provider interface.
type StackOverflowProvider struct {
}

// BuildURI generates a search URL for Stack Overflow.
func (p *StackOverflowProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://stackoverflow.com/search?q=%s", url.QueryEscape(q))
}
