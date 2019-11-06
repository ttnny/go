package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Registering a Request Handler
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)

		// To read a GET
		r.URL.Query().Get("token")

		// To POST a field from HTML form
		r.FormValue("email")

		// Serving static assets
		fs := http.FileServer(http.Dir("static/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))
	})

	// Listen for HTTP Connections
	http.ListenAndServe(":80", nil)
}
