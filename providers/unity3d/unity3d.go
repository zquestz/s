package unity3d

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("unity3d", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for Unity3D.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://unity3d.com/search?gq=%s", url.QueryEscape(q))
}
