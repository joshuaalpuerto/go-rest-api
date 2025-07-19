package main

import (
	"fmt"
	"net/http"

	"github.com/joshuaalpuerto/go-rest-api/config"
)

func main() {
	conf := config.New()
	host := conf.Server.Host
	port := conf.Server.Port

	// The HandleFunc function registers a handler function for a given pattern.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// The Fprintf function formats according to a format specifier and writes to w.
		fmt.Fprintf(w, "Hello, World!")

	})

	fmt.Println("Server is starting on port 8080...")
	// ListenAndServe starts an HTTP server with a given address and handler.
	// It blocks until the server is shut down.
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
