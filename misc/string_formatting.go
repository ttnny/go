package main

import "fmt"

func main() {
	brand := "Google"

	fmt.Printf("%q\n", brand) 	// "Google"
	fmt.Printf("%s\n", brand)	// Google
	fmt.Printf("%v\n", brand)	// Google
	fmt.Printf("%d\n", brand)	// Google

	fmt.Printf("RGB: %d, %d, and %d\n", 70, 215, 140)
	fmt.Printf("RGB: %v, %v, and %v\n", 70, 215, 140)
	fmt.Printf("RGB: %[2]v, %[3]v, and %[1]v\n", 70, 215, 140)

	fmt.Printf("%T\n", brand)	// string

	num := 364.907

	fmt.Printf("Days: %f\n", num)   // 364.907000
	fmt.Printf("Days: %.1f\n", num) // 364.9
	fmt.Printf("Days: %.0f\n", num) // 365
}