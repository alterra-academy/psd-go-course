package main

import "fmt"

func main() {
	// long declaration
	var even_numbers []int
	fmt.Printf("elements = %v, len = %d, cap = %d\n", even_numbers, len(even_numbers), cap(even_numbers))

	// long declaration with values
	var odd_numbers = []int{1, 3, 5, 7, 9}
	fmt.Printf("elements = %v, len = %d, cap = %d\n", odd_numbers, len(odd_numbers), cap(odd_numbers))

	// short declaration with values
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("elements = %v, len = %d, cap = %d\n", numbers, len(numbers), cap(numbers))

	// using make function
	var primes = make([]int, 5, 10)
	fmt.Printf("elements = %v, len = %d, cap = %d\n", primes, len(primes), cap(primes))
}