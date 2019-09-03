package main

import (
	"fmt"
)

func main() {
	// movies := [3]string{"Avengers", "Iron Man", "Spider Man"}
	movies := [...]string{"Iron Man"}
	movies = append(movies, "OK")

	movies[0] = "Now you see me"
	fmt.Printf("Movies: %v\n", movies)

	// Get the length of array
	fmt.Printf("Total movies: %v\n", len(movies))

	//Copy of array
	newMovies := movies
	newMovies[0] = "Avengers"
	fmt.Printf("New movies: %v\n", newMovies)

	//Pointer in array, ll cover pointers in details later
	oldMovies := &movies
	oldMovies[0] = "Harry potter"
	fmt.Printf("Movies: %v\n", movies)
	fmt.Printf("O ld movies: %v\n", oldMovies)
}
