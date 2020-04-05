package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Always use Join() instead of concatenating /s or \s manually
	path := filepath.Join("dir", "sub-dir", "filename")
	fmt.Println("Path: ", path)

	// Join() also removes superfluous separators and directory changes
	fmt.Println(filepath.Join("dir//", "filename"))
	fmt.Println(filepath.Join("dir/../dir", "filename"))

	// Dir() splits a path to the directory
	fmt.Println("Dir:", filepath.Dir(path))

	// Base() splits a path to the file
	fmt.Println("Base:", filepath.Base(path))

	// IsAbs() checks whether a path is absolute
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	// Ext() splits the extension out of filename
	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// TrimSuffix() returns the filename with extension removed
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel() finds relative path between a base and a target
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	// Environment variables

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}