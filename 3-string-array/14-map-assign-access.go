package main

import "fmt"

func main() {
	// long declaration with value
	var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
	fmt.Println(salary_a, len(salary_a))

	salary_a["nabilah"] = 7000 // assign value
	fmt.Println(salary_a)

	delete(salary_a, "iswanul") // remove value by key
	fmt.Println(salary_a)

	value, exist := salary_a["umam"] // check key is exist
	fmt.Println(value, exist)

	for key, value := range salary_a { // loop over maps
		fmt.Println("->", key, value)
	}
}
