package main

import (
	"fmt"
)

const (
	// errorA = iota // so we can use index 1 values and let 0th one as an error
	// best practice just ignore the value and dont assign in memory
	// _ = iota

	// OR just shipt the value.
	a = 1 << iota
	// b = iota
	// c = iota
	// We can like this
	b
	c
)

func main() {
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
}
