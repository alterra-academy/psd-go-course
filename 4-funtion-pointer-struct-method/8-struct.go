package main

import "fmt"

type Person struct {
  FirstName string
  LastName  string
  Age       int
}

func main() {
	// long declaration
	var Person0 Person
	Person0.FirstName = "Muchson"
	Person0.LastName = "Ibi"
	Person0.Age = 27
	fmt.Println(Person0.FirstName, Person0.LastName, Person0.Age)

	// long declaration with assigned value
	var Person1 = Person{"Rizky", "Kurniawan", 26}
	fmt.Println(Person1)

	// long declaration with assigned value each name fields
	var Person2 = Person{
		FirstName: "Iswanul",
		LastName:  "Umam",
		Age:       25,
	}
	fmt.Println(Person2)

	// sort declaration
	Person3 := Person{"Pranadya", "Bagus", 23}
	fmt.Println(Person3)

	// short declaration with new keyword
	Person4 := new(Person)
	Person4.FirstName = "Muhammad"
	Person4.LastName = "Ismail"
	Person4.Age = 30
	
	fmt.Println(*Person4)
}
