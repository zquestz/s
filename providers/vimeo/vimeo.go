package vimeo

import (
	"fmt"
	"net/url"
	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("vimeo", &VimeoProvider{})
}

//VimeoProvider adheres to the Provider interface.
type VimeoProvider struct {
}

//BuildURI generates a search URL for Vimeo.
func (p *VimeoProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://vimeo.com/search?q=%s", url.QueryEscape(q))
}
