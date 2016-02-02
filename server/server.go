package server

import (
	"fmt"
	"net/http"
)

// Run sets up and starts the http server.
func Run(port int, provider string) error {
	err := validatePort(port)
	if err != nil {
		return err
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index(provider, w, r)
	})
	http.HandleFunc("/search", search)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		return err
	}

	return nil
}

func validatePort(port int) error {
	if port < 1 || port > 65535 {
		return fmt.Errorf("Invalid port requested. Valid values are 1-65535.")
	}

	return nil
}
