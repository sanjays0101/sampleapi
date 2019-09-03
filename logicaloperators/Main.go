package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	if a > 9 && b > 10 {
		fmt.Println(true)
	}

	var c = a > 9 && a == 10
	fmt.Println(c)

	var d = a > 10 || a == 10
	fmt.Println(d)
}
