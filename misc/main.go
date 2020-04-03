package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := cards.deal(3)

	fmt.Println(hand, remainingDeck)
}
