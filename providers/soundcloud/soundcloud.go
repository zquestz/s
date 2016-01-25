package soundcloud

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("soundcloud", &SoundCloudProvider{})
}

// SoundCloudProvider adheres to the Provider interface.
type SoundCloudProvider struct {
}

// BuildURI generates a search URL for SoundCloud.
func (p *SoundCloudProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://soundcloud.com/search?q=%s", url.QueryEscape(q))
}
