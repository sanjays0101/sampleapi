package main

import "fmt"

func modifyMyHeight(youHeight *int8) {
	*youHeight = *youHeight + 1
	fmt.Printf("Modified height: %v\n", *youHeight)
}

func main() {

	var height int8 = 20
	fmt.Printf("Original height: %v\n", height)
	modifyMyHeight(&height)
	fmt.Printf("Original height: %v\n", height)

}
