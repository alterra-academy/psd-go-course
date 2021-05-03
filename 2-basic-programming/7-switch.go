package main

import "fmt"

func main() {
	var today int = 2  
	switch today {
		case 1:
			fmt.Printf("Today is Monday")
		case 2:
			fmt.Printf("Today is Tuesday")
		default:
			fmt.Printf("Invalid Date")            
	}	
}