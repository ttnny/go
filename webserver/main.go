package webserver

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
	})

	err := http.ListenAndServe(":8082", nil)

	if err != nil {
		panic(err)
	}
}