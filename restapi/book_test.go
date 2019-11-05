package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Chicken Teriyaki", Author: "C.K.", ISBN: "0123456789"}
	json := book.ToJSON()

	assert.Equal(t, `{"Title":"Chicken Teriyaki","Author":"C.K.","ISBN":"0123456789"}`, string(json), "Book JSON marshalling wrong.")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"Title":"Chicken Teriyaki", "Author":"C.K.","ISBN":"0123456789"}`)
	book := FromJSON(json)

	assert.Equal(t, Book{Title: "Chicken Teriyaki", Author: "C.K.", ISBN: "0123456789"}, book, "Book JSON unmarshalling wrong.")
}
