package main

import "fmt"

func main() {
	var primes []int
	fmt.Printf("s = %v, len = %d, cap = %d\n", primes, len(primes), cap(primes))

	if primes == nil {
		fmt.Println("s is nil")
	}
}
