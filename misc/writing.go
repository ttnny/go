package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	d1 := []byte("hello\ngo\n")

	// Write to a new file and handle errors.
	err := ioutil.WriteFile("misc/sample1.txt", d1, 0644)
	if err != nil {	}

	// Write to an existing file and handle errors.
	file, err := os.Create("misc/sample1.txt")
	if err != nil {}
	defer file.Close()

	n2, err := file.Write(d1)
	if err != nil {}
	fmt.Printf("Wrote %d bytes\n", n2)
}