package main

import "fmt"

// 3rd
// Package level declaration
var j int = 30

// 3rd End

//4th
// Declare var in bulk
var (
	actorName       string  = "Robert Downey"
	favMovie        string  = "Iron Man"
	movieRating     float32 = 4.9 // out of 5
	totalCharacters int     = 100
)

//4th End

func main() {

	// 1st
	// var i int
	// i = 10
	// i = 20
	/*
		Another way to declare variable
		1. var i int  = 10
		Or with short hand
		2.  i := 10
	*/
	// fmt.Println(i)
	// 1st End

	// 2nd
	fmt.Printf("%v, %T", j, j)

	// 2nd End

}
