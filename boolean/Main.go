package main

import "fmt"

func main() {
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	n1 := 1 == 1
	n2 := 1 == 2
	fmt.Printf("%v, %T\n", n1, n1)
	fmt.Printf("%v, %T\n", n2, n2)
	fmt.Printf("Test_101")
}
