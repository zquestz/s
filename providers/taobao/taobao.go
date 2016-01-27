package taobao

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("taobao", &TaobaoProvider{})
}

// TaobaoProvider adheres to the Provider interface.
type TaobaoProvider struct {
}

// BuildURI generates a search URL for Taobao.
func (p *TaobaoProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://s.taobao.com/search?q=%s", url.QueryEscape(q))
}
