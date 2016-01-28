package codepen

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("codepen", &CodePenProvider{})
}

//CodePenProvider adheres to the Provider interface.
type CodePenProvider struct {
}

//BuildURI generates a search URL for CodePen.
func (p *CodePenProvider) BuildURI(q string) string {
	return fmt.Sprintf("http://codepen.io/search/pens?q=%s", url.QueryEscape(q))
}
