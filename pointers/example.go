package main

import "fmt"

func main() {
	a := 1
	var pa *int

	fmt.Printf("Pointer 'pa' of type %T with value %v\n", pa, pa)

	pa = &a
	fmt.Printf("Pointer 'pa' of type [%T] at memory address [%v] with value [%v]\n", pa, pa, *pa)

	// Another example with func new(Type) *Type
	// new() allocates a new memory address and return a pointer pointing to it
	pb := new(string)
	fmt.Printf("Pointer 'pb' of type [%T] at memory address [%v] with value [%v]\n", pb, pb, *pb)
}