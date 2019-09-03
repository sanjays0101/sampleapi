package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}
	movies := []string{"Iron Man", "Avengers", "Spider Man", "Hulk"}
	expLastMovie := movies[:len(movies)-1]
	fmt.Printf("Last movie: %v\n", expLastMovie)
	// fmt.Printf("Movies: %v\n", movies)

	//Remove middle item from the movies slice
	/*
		n = 100
		rm = 78
		var xyz := append(n[:77], n[78:]...)
	*/
	rmmMovies := append(movies[:2], movies[3:]...)
	//
	fmt.Printf("Movies after mid val rem: %v\n", movies)
	// test := movies[:1]
	fmt.Printf("rmmMovies: %v\n", rmmMovies)

	// Remove first value from the slice
	a1 := a[1:]
	fmt.Printf("a1: %v\n", a1)

	// Remove last value from slice
	b := a[:len(a)-1]
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	// Remove middle value from slice
	d := append(a[:1], a[3:]...)
	fmt.Printf("d: %v\n", d)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	//fmt.Printf("Mid value: %v\n", b[len(nums)/2])
	//midVal := append(nums[:len(nums/2) - 1])
	fmt.Printf("%v\n", len(nums)/2-1)
	fmt.Printf("Mid val:%v\n", nums[len(nums)/2-1])
	onlyMidVal := nums[len(nums)/2-1:]
	fmt.Printf("Only mid val: %v\n", onlyMidVal)
	fmt.Print
}
