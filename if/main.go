package main

import "fmt"

func main() {
	var num int = 10

	if num > 9 {
		fmt.Println(num)
	}

	if true {
		fmt.Println("Always print.")
	}

	if false {
		fmt.Println("Never print this line")
	}

	statePopulation := map[string]int{
		"KA": 10000,
		"MH": 100000,
	}
	if pop, ok := statePopulation["KA"]; ok {
		fmt.Println(pop)
	}
}
