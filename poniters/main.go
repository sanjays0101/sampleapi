package main

import "fmt"

func printAge(yourAge *int8) int8 {

	*yourAge = *yourAge + 1
	return *yourAge
}

func main() {

	var age int8 = 20

	fmt.Printf("Original age: %v\n", age)
	fmt.Println(printAge(&age))
	fmt.Printf("Original  age: %v\n", age)
	//printAge(&age)

}
