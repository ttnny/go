package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.Open("misc/sample.txt")
	if err != nil {
		// Handle errors.

		// Open opens the named file for reading. If successful, methods on
		// the returned file can be used for reading; the associated file
		// descriptor has mode O_RDONLY.
		// If there is an error, it will be of type *PathError.
	}
	defer file.Close()

	fiveChars := make([]byte, 5)

	// Read the first 5 chars to the byte slice 'fiveChars'
	// and return the number of bytes read and any error encountered.
	// At end of file, Read returns 0, io.EOF
	n, err := file.Read(fiveChars)
	fmt.Printf("%d bytes: %s\n", n, string(fiveChars))

	// Seek to a known location and Read from there.
	threeChars := make([]byte, 3)

	// Set the offset for the next Read or Write on file
	o, err := file.Seek(2, 0)

	// Read 2 chars starting from index 2 as set above
	n1, err := file.Read(threeChars)
	fmt.Printf("%d bytes @ %d: %v\n", n1, o, string(threeChars))



	o1, err := file.Seek(2, 0)
	fourChars := make([]byte, 4)

	// ReadAtLeast reads from r into buf until it has read at least min bytes.
	// It returns the number of bytes copied and an error if fewer bytes were read.
	// The error is EOF only if no bytes were read.
	n2, err := io.ReadAtLeast(file, fourChars, 4)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o1, string(fourChars))


	// Rewind with Seek(0, 0)
	_, err = file.Seek(0, 0)
	check(err)

	// Implement a buffered reader.
	// Peek returns the next n bytes without advancing the reader.
	r4 := bufio.NewReader(file)
	b4, err := r4.Peek(6)

	fmt.Printf("5 bytes: %s\n", string(b4))
}