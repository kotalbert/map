package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println(colors)

	var emptyMap map[string]string
	fmt.Println(emptyMap)

	colors["black"] = "#000000"
	colors["white"] = "#ffffff"
	colors["bogus"] = "123"

	fmt.Println(colors)
	delete(colors, "bogus")
	fmt.Println(colors)

	b := colors["black"]
	fmt.Println(b)

}
