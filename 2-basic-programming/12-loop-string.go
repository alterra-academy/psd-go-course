package main

import "fmt"

func main() {
	sentence := "Hello";
	
	for i := 0; i < len(sentence); i++ {
		fmt.Println(string(sentence[i]))
	}
	
	fmt.Println("-----")
	
	for pos, char := range sentence {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	} 
}

