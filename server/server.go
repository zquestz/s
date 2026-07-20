package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NYTimes/gziphandler"
)

// Run sets up and starts the http server.
// It requires a valid port to bind to, and the
// default provider to use for web searches.
func Run(port int, cert string, key string, provider string, verbose bool) error {
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
		search(provider, verbose, w, r)
	})

	http.Handle("/", gziphandler.GzipHandler(indexHandler))
	http.Handle("/search", gziphandler.GzipHandler(searchHandler))

	setupFaviconHandlers()

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	if cert != "" && key != "" {
		err = srv.ListenAndServeTLS(cert, key)
	} else {
		err = srv.ListenAndServe()
	}

	if err != nil {
		return fmt.Errorf("HTTP Server: %s", err)
	}

	return nil
}

func validatePort(port int) error {
	if port < 1 || port > 65535 {
		return fmt.Errorf("invalid port requested, valid values are 1-65535")
	}

	return nil
}
