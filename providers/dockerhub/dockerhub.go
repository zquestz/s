package dockerhub

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("dockerhub", &Provider{})
}

type Provider struct{}

// BuildURI generates a search URL for Docker Hub.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("https://hub.docker.com/search/?q=%s", url.QueryEscape(q))
}
