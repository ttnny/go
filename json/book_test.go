package main

import (
	"github.com/stretchr/testify"
	"testing"
)

func TestToJSON(t *testing.T) {
	book := Book {Title: "Chicken Teriyaki", Author: "C.K.", ISBN: "0123456789"}
	json := book.ToJSON()

	assert.Equal(t, `{"Title":"Chicken Teriyaki", "Author":"C.K.","ISBN":"0123456789"}`, string(json), "Book JSON marshalling wrong.")
}

func TestFromJSON(t *testing.T) {
	assert.True(t, true, "Implement me.")
}
