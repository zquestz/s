package coursera

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("coursera", &Coursera{})
}

// Coursera adheres to the Provider interface.
type Coursera struct {
}

// BuildURI generates a search URL for Coursera.
func (p *Coursera) BuildURI(q string) string {
	return fmt.Sprintf("https://www.coursera.org/courses/?query=%s", url.QueryEscape(q))
}
