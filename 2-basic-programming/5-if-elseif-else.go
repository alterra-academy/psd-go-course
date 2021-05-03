package main

import "fmt"

func main() {
	hour := 15
	if hour < 12 {
		fmt.Println("Selamat Pagi")
	} else if hour < 18 {
		fmt.Println("Selamat Sore")
	} else {
		fmt.Println("Selamat Malam")
	}
}