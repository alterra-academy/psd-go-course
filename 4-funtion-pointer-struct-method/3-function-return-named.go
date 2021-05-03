package main

import "fmt"

// function having named return parameter
func multiplication(a, b int) (mul int) {
	mul = a * b
	return
}

func main() {
	m := multiplication(5, 5)
	fmt.Println("5 x 5 = ", m)
}
