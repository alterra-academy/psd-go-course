package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Obtaining a slice from an array `array`
	// array[low:high]
	var primes = [5]int{2, 3, 5, 7, 11}

	// Creating a slice from the array
	var part_primes []int = primes[1:4]

	// part_primes = append(part_primes, 10000)
	// menambah data ke slice akan menambah data ke array juga

	fmt.Println(reflect.ValueOf(part_primes).Kind())
	fmt.Println(part_primes)
}
