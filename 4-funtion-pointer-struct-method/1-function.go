package main

import "fmt"

// Without parameter
func sayHello() {
	fmt.Println("Hello")
}

// With parameter
func greeting(hour int) {
	if hour < 12 {
		fmt.Println("Selamat Pagi")
	} else if hour < 18 {
		fmt.Println("Selamat Sore")
	} else {
		fmt.Println("Selamat Malam")
	}
}

func main() {
	hour := 15
	greeting(hour)
}
