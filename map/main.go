package main

import "fmt"

func main() {
	//Declare Empty Map:
	colors2 := make(map[int]string)
	colors2[1] = "One"

	// Decalare + Initial Map
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
	}

	fmt.Println(colors)
	fmt.Println(colors2)
}
