package main

import "fmt"

func main() {
	//Declare Empty Map:
	//colors2 := make(map[int]string)
	//colors2[1] = "One"

	// Declare + Initial Map
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
