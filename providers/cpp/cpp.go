package cpp

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("cpp", &CppProvider{})
}

// CppProvider adheres to the Provider interface.
type CppProvider struct {
}

// BuildURI generates a search URL for Cpp.
func (p *CppProvider) BuildURI(q string) string {
	return fmt.Sprintf("http://www.cplusplus.com/search.do?q=%s", url.QueryEscape(q))
}
