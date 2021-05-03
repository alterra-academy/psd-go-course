package main

import "fmt"

type Person struct {
	name string // Both non exported fields.
	age  int
}

func (P Person) GetName() string {
	return P.name + " amazing!"
}

func (P *Person) IncreaseAge() {
	P.age = P.age + 1
}

func main() {
	PersonA := Person{"John", 50}
	fmt.Printf("%v\n", PersonA)
	fmt.Println(PersonA.GetName())

	PersonA.IncreaseAge()
	fmt.Println(PersonA.age)
}
