package main

import (
	"fmt"
)

func main() {
	var size = new(int)
	fmt.Printf("Size value is %d \n", *size)
	fmt.Printf("Type is %T \n", size)
	fmt.Printf("Address is %v \n", size)
	*size = 85
	fmt.Println("New size value is", *size)
}
