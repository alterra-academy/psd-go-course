package main

import "fmt"

func main() {
	var today int = 2  
	switch {
	   case today == 1:
		   fmt.Printf("Today is Monday")
	   case today == 2:
		   fmt.Printf("Today is Tuesday")
	   default:
		   fmt.Printf("Invalid Date")            
	}	
}