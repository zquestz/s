package dockerhub

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("dockerhub", &DockerHubProvider{})
}

// DockerHubProvider adheres to the Provider interface.
type DockerHubProvider struct {
}

// BuildURI generates a search URL for Docker Hub.
func (p *DockerHubProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://hub.docker.com/search/?q=%s", url.QueryEscape(q))
}
