package main

import (
	"fmt"
	"reflect"
)

func main() {
	primes := [...]int{2, 3, 3}

	fmt.Println(reflect.ValueOf(primes).Kind())
	fmt.Println(len(primes))
}
