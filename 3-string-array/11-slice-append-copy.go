package main

import "fmt"

func main() {
	var colors = []string{"red", "green", "yellow"}
	colors = append(colors, "purple")

	copied_colors := make([]string, 10)

	copy(copied_colors, colors) // copy from colors to copied_colors
	fmt.Println(copied_colors)
}