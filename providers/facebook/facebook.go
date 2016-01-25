package facebook

import (
  "fmt"
  "net/url"

  "github.com/zquestz/s/providers"
)

func init() {
  providers.AddProvider("facebook", &FacebookProvider{})
}

// FacebookProvider adheres to the Provider interface.
type FacebookProvider struct {
}

// BuildURI generates a search URL for facebook.
func (p *FacebookProvider) BuildURI(q string) string {
  return fmt.Sprintf("https://www.facebook.com/search/top/?q=%s", url.QueryEscape(q))
}
