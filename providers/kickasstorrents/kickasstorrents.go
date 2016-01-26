package kickasstorrents

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("kickasstorrents", &KickAssTorrentsProvider{})
}

// QuoraProvider adheres to the Provider interface.
type KickAssTorrentsProvider struct {
}

// BuildURI generates a search URL for KickAss Torrents.
func (p *KickAssTorrentsProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://kat.cr/usearch/%s/", url.QueryEscape(q))
}
