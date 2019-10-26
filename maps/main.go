package main

import "fmt"

func main() {
	colors := map[string]string {
		"red":   "#ff0000",
		"green": "#008000",
		"blue":  "#0000ff",
	}

	colors["white"] = "#ffffff"
	delete(colors, "green")

	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println()
	}
}