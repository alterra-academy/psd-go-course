package main

import "fmt"

func main() {
	var countries [2]string

	countries[0] = "India"  // Assign a value to the first element
	countries[1] = "Canada" // Assign a value to the second element

	fmt.Println(countries[0]) // Access the first element value
	fmt.Println(countries[1]) // Access the second element value
}