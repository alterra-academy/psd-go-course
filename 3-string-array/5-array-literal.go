package main

import "fmt"

func main() {
	odd_numbers := [5]int{1, 3, 5, 7, 9}   // Intialized with values
	var even_numbers [5]int = [5]int{0, 2, 4} // Partial assignment

	fmt.Println(odd_numbers)
	fmt.Println(even_numbers)
}