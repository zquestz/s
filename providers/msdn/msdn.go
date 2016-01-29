package msdn

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("msdn", &MSDNProvider{})
}

//MSDNProvider adheres to the Provider interface.
type MSDNProvider struct {
}

//BuildURI generates a search URL for MSDN.
func (p *MSDNProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://social.msdn.microsoft.com/Search/en-US?query=%s", url.QueryEscape(q))
}
