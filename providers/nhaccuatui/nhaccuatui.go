package nhaccuatui

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("nhaccuatui", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for Nhaccuatui.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("http://www.nhaccuatui.com/tim-kiem?q=%s", url.QueryEscape(q))
}
