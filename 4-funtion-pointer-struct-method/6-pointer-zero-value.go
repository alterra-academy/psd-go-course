package main

import (
  "fmt"
)

func main() {
	number_a := 25
	var number_b *int
	
	if number_b == nil {
		fmt.Println("number_b is", number_b)
		number_b = &number_a
		fmt.Println("number_b after init : is", *number_b)
	}
}
