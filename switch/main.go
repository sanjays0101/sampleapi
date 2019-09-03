package main

import "fmt"

func main() {
	var age int8 = 20

	switch age {
	case 0:
		fmt.Println("Nope")
	case 20:
		fmt.Println("Yes")
	default:
		fmt.Println("Default")
	}

	// some advance

	switch i := 2 + 3; i {
	case 5:
		{
			fmt.Println("Five")
		}
	}
}
