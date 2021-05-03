package main

import "fmt"

func main() {
	// long declaration
	var salary = map[string]int{}
	fmt.Println(salary)

	// long declaration with value
	var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
	fmt.Println(salary_a)

	// short declaration
	salary_b := map[string]int{}
	fmt.Println(salary_b)

	// using make
	var salary_c = make(map[string]int)
	salary_c["doe"] = 7000 // assign value
	fmt.Println(salary_c)
}


