package main

import "fmt"

type bot interface {
	getGreeting() string
}

type empty interface{} // or just: var any interface{}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	// Try empty interface
	var any empty
	any = []int {1, 2, 3}
	any = map[int]bool{1: true, 2: false}
	any = "Hello"
	any = 3

	// any = any * 2 <-- this won't compile, need type assertion to extract number
	any = any.(int) * 2
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
