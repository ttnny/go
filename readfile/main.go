package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Method 1: Read the while file
	token1, err := ioutil.ReadFile("readfile/token.txt")
	check(err)
	fmt.Println(string(token1))

	// Method 2: Open the file, then read the first 5 bytes
	token2, err := os.Open("readfile/token.txt")
	check(err)

	b2 := make([]byte, 5)
	n2, err := token2.Read(b2)
	check(err)
	fmt.Printf("%d bytes: %s\n", n2, string(b2[:n2]))

	_ = token2.Close()

	// Method 3: Open the file, then seek to a known location in the file and read from there
	token3, err := os.Open("readfile/token.txt")
	check(err)

	o3, err := token3.Seek(6, 0)
	check(err)

	b3 := make([]byte, 2)
	n3, err := token3.Read(b3)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n3, o3)
	fmt.Printf("%v\n", string(b3[:n3]))

	_ = token3.Close()

	// Method 4: Same as method 3 but using io.ReadAtLeast()
	token4, err := os.Open("readfile/token.txt")
	check(err)

	o4, err := token4.Seek(6, 0) // Seek(0, 0) to rewind
	check(err)

	b4 := make([]byte, 2)
	n4, err := io.ReadAtLeast(token4, b4, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n4, o4, string(b4))

	_ = token4.Close()

	// Method 5: Same as method 2 but using a buffered reader to read
	token5, err := os.Open("readfile/token.txt")
	check(err)

	r5 := bufio.NewReader(token5)
	b5, err := r5.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b5))

	_ = token5.Close()
}
