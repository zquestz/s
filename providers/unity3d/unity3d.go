package unity3d

import (
	"fmt"
	"net/url"
	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("unity3d", &Unity3DProvider{})
}

//Unity3DProvider adheres to the Provider interface.
type Unity3DProvider struct {
}

//BuildURI generates a search URL for Unity3D.
func (p *Unity3DProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://unity3d.com/search?gq=%s", url.QueryEscape(q))
}
