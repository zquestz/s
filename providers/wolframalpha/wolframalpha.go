package wolframalpha

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("wolframalpha", &WolframAlphaProvider{})
}

// WolframAlphaProvider adheres to the Provider interface.
type WolframAlphaProvider struct {
}

// BuildURI generates a search URL for WolframAlpha.
func (p *WolframAlphaProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.wolframalpha.com/input/?i=%s", url.QueryEscape(q))
}
