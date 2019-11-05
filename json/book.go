package main

import (
	"encoding/json"
	"net/http"
)

// Book type with Name, Author, and ISBN
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// ToJSON to be used for marshalling of Book type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.MarshalIndent(b, "", "")

	if err != nil {
		panic(err)
	}

	return ToJSON
}

// FromJSON to be used for unmarshalling of Book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)

	if err != nil {
		panic(err)
	}

	return book
}

var Books = []Book {
	Book{"Book Title 1", "C.K.1", "0123456789"},
	Book{"Book Title 2", "C.K.2", "9876543210"},
}

// BooksHandleFunc to be used as http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err:= json.MarshalIndent(Books, "","")

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset-utf-8")
	w.Write(b)
}
