package baidu

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("baidu", &BaiduProvider{})
}

// BaiduProvider adheres to the Provider interface.
type BaiduProvider struct {
}

// BuildURI generates a search URL for Google.
func (p *BaiduProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.baidu.com/s?wd=%s", url.QueryEscape(q))
}
