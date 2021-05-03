package main

import "fmt"

func main() {
	var name string = "John"
	var nameAddress *string = &name
	fmt.Println("name (value)   :", name)                // John
	fmt.Println("name (address) :", &name)               // 0xc000010050
	fmt.Println("nameAddress (value)   :", *nameAddress) // John
	fmt.Println("nameAddress (address) :", nameAddress)  // 0xc000010050
}