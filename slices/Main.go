package main

import (
	"fmt"
)

func main() {
	movies := []string{"Iron Man", "Spider Man", "Avengers"}
	fmt.Printf("Movies: %v\n", movies)
	fmt.Printf("Total movies: %v\n", len(movies))

	// Capacity in slice
	fmt.Printf("Total movies %v\n", cap(movies))

	// Slices points to the address or creates new copy ?
	newMovies := movies
	newMovies[0] = "Now you see me"
	fmt.Printf("Movies:  %v\n", movies)
	fmt.Printf("New Movies: %v\n", newMovies)

	// More on slice
	a := []int{1, 2, 3, 4, 5}
	b := a[:] // slice of all elements
	c := a[3:] // slice off up to 3
	d := a[2:4] // slice off up to 2 and print up to 4 from beggining
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("C: %v\n", c)
	fmt.Printf("d: %v\n", d)

	//Append values to the slice
	marks := []int{}
	marks = append(marks, 1, 2, 3, 4, 5, 6)
	marks = append(marks, 7, 8)
	fmt.Printf("Marks: %v\n", marks)

	// with length and capacity
	numbers := make([]int, 0, 5)
	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14)
	fmt.Printf("Numbers: %v\n", numbers)

}
