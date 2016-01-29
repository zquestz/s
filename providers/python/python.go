package python

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("python", &PythonProvider{})
}

//PythonProvider adheres to the Provider interface.
type PythonProvider struct {
}

//BuildURI generates a search URL for Python.
func (p *PythonProvider) BuildURI(q string) string {
	return fmt.Sprintf("https://www.python.org/search/?q=%s", url.QueryEscape(q))
}
