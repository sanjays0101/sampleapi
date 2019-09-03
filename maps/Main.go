package main

import "fmt"

func main() {
	// map syntax: map[key_type]value_type{}
	statePopulation := map[string]int{
		"Karnataka":   61095297,
		"Maharashtra": 112374333,
	}
	fmt.Println(statePopulation)

	// Second way

	countryPopulation := make(map[string]int)
	countryPopulation = map[string]int{
		"India": 10,
		"USA":   20,
	}
	countryPopulation["UK"] = 5
	fmt.Println(countryPopulation)

	//delete entry from map
	delete(countryPopulation, "UK")
	fmt.Println(countryPopulation)

	// Lets make mistakes
	fmt.Println(countryPopulation["UK"]) // result 0 ? no one in UK ?

	// Fix it
	UKPopulation, ok := countryPopulation["UK"]
	fmt.Println(UKPopulation, ok)

	fmt.Printf("Total states:  %v\n", len(statePopulation))

}
