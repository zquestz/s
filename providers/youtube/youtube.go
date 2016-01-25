package youtube

import (
  "fmt"
  "net/url"

  "github.com/zquestz/s/providers"
)

func init() {
  providers.AddProvider("youtube", &YouTubeProvider{})
}

// YouTubeProvider adheres to the Provider interface.
type YouTubeProvider struct {
}

// BuildURI generates a search URL for YouTube.
func (p *YouTubeProvider) BuildURI(q string) string {
  return fmt.Sprintf("https://www.youtube.com/results?search_query=%s", url.QueryEscape(q))
}
