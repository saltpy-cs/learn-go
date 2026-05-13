package main

import "fmt"

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Printf("%s: %s\n", k, v)
	}
}

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	printMap(colors)
}
