package nhaccuatui

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("nhaccuatui", &NhaccuatuiProvider{})
}

// NhaccuatuiProvider interface.
type NhaccuatuiProvider struct {
}

// BuildURI generates a search URL for NhaccuatuiProvider.
func (p *NhaccuatuiProvider) BuildURI(q string) string {
	return fmt.Sprintf("http://www.nhaccuatui.com/tim-kiem?q=%s", url.QueryEscape(q))
}
