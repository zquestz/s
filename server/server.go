package server

import (
	"fmt"
	"net/http"

	"github.com/NYTimes/gziphandler"
)

// Run sets up and starts the http server.
// It requires a valid port to bind to, and the
// default provider to use for web searches.
func Run(port int, cert string, key string, provider string) error {
	err := validatePort(port)
	if err != nil {
		return err
	}

	indexHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		index(provider, w, r)
	})

	searchHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		search(provider, w, r)
	})

	http.Handle("/", gziphandler.GzipHandler(indexHandler))
	http.Handle("/search", gziphandler.GzipHandler(searchHandler))

	if cert != "" && key != "" {
		err = http.ListenAndServeTLS(fmt.Sprintf(":%d", port), cert, key, nil)
		if err != nil {
			return fmt.Errorf("HTTP Server: %s", err)
		}
	} else {
		err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			return fmt.Errorf("HTTP Server: %s", err)
		}
	}

	return nil
}

func validatePort(port int) error {
	if port < 1 || port > 65535 {
		return fmt.Errorf("Invalid port requested. Valid values are 1-65535.")
	}

	return nil
}
