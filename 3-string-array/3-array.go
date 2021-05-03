package main

import (
	"fmt"
	"reflect"
)
  
func main() {
	var primes [5]int
	var countries [5]string
  
	fmt.Println(reflect.ValueOf(primes).Kind())
	fmt.Println(reflect.ValueOf(countries).Kind())
}
  
  
  