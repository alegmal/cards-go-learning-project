package main

import (
	"fmt"
)

func main()  {
	colors := map[string]string {
		"red": "#1234",
		"white": "#5533",
		"yellow": "#7676",
	}
	
	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m{
		fmt.Println(key, value)
	}
}