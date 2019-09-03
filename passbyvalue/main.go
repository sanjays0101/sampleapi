package main

import "fmt"

func PrintAge(yourAge int8) {
	fmt.Println(yourAge)
}

func IncMyAge(yourAge int8) {
	yourAge = yourAge + 1
	fmt.Println(yourAge)
}

func modifyMyHeight(youHeight int8) int8 {
	return youHeight + 1
}

func main() {

	var age int8 = 20
	PrintAge(age)

	// Modify value
	IncMyAge(age)

	// Modify source value ?

	var height int8 = 20
	fmt.Printf("Original height: %v\n", height)

	height = modifyMyHeight(height)
	fmt.Printf("Modified height: %v\n", height)

}
