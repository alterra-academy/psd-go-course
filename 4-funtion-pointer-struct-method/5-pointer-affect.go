package main

import "fmt"

func main() {
	var name string = "John"
	var nameAddress *string = &name
	fmt.Println("name (value)   :", name) // John
	fmt.Println("name (address) :", &name) // 0xc20800a220
	fmt.Println("nameAddress (value)   :", *nameAddress) // John
	fmt.Println("nameAddress (address) :", nameAddress) // 0xc20800a220

	name = "Doe"

	fmt.Println("")  
	fmt.Println("name (value)   :", name) // Doe
	fmt.Println("name (address) :", &name) // 0xc20800a220
	fmt.Println("nameAddress (value)   :", *nameAddress) // Doe
	fmt.Println("nameAddress (address) :", nameAddress) // 0xc20800a220
}