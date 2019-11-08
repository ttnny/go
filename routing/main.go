package routing

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Create a new request router
	r := mux.NewRouter()

	// Register a request handler
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		// Get data from the extracted segments {title} and {page}
		// mux.Vars(r) takes the http.Request as parameter and returns a map of the segments
		vars := mux.Vars(r)

		title := vars["title"] // the book title slug
		page := vars["page"]   // the page

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	// Setting the HTTP server's router
	// With gorilla/mux, use 'r' as our main router instead
	// of 'nil' (a default router from the net/http package)
	http.ListenAndServe(":80", r)
}

func otherMuxFeatures() {
	// Create a new request router
	r := mux.NewRouter()

	// *** Methods ***
	// Restrict the request handler to specific HTTP methods
	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	// *** Hostnames & Subdomains ***
	// Restrict the request handler to specific hostnames or subdomains
	r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

	// *** Schemes ***
	// Restrict the request handler to http/https
	r.HandleFunc("/secure", SecureHandler).Schemes("https")
	r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	// *** Path Prefixes & Subrouters ***
	// Restrict the request handler to specific path prefixes
	bookrouter := r.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/", AllBooks)
	bookrouter.HandleFunc("/{title}", GetBook)
}

// Just dummy data
func CreateBook(writer http.ResponseWriter, request *http.Request) {}
func ReadBook(writer http.ResponseWriter, request *http.Request) {}
func UpdateBook(writer http.ResponseWriter, request *http.Request) {}
func DeleteBook(writer http.ResponseWriter, request *http.Request) {}
func BookHandler(writer http.ResponseWriter, request *http.Request) {}
func SecureHandler(writer http.ResponseWriter, request *http.Request) {}
func InsecureHandler(writer http.ResponseWriter, request *http.Request) {}
func AllBooks(writer http.ResponseWriter, request *http.Request) {}
func GetBook(writer http.ResponseWriter, request *http.Request) {}