package main

import (
	"fmt"
	"math"
)

const t1 = "top 1"
const t2 string = "top 2"

func main() {
	const n = 500000000
	const d = 3e20 / n

	fmt.Println(t1) // top 1
	fmt.Println(t2) // top 2

	fmt.Println(d)           // 6e+11
	fmt.Println(int64(d))    // 600000000000
	fmt.Println(math.Sin(n)) // -0.28470407323754404
}
