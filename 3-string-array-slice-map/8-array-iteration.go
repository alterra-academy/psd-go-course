package main

import "fmt"

func main() {
	primes := [5]int{2, 3, 5}

	// technique 1
	for index := 0; index < len(primes); index++ {
	  fmt.Println(primes[index])
	}
	// technique 2
	for index, element := range primes {
	  fmt.Println(index, "=>", element)
	}
	for _, value := range primes {
	  fmt.Println(value)
	}
	// technique 3
	index := 0
	for range primes {
	  fmt.Println(primes[index])
	  index++
	}	
}