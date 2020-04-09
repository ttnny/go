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



	o3, err := file.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(file, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = file.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(file)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
}