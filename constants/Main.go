package main

import (
	"fmt"
)

func main() {
	const myConst int = 10
	fmt.Printf("%v, %T\n", myConst, myConst)

	// 2nd
	// const a int = 10
	// const b int16 = 20
	// fmt.Printf("%v, %T\n", a+b, a)

	// 3rd
	const a = 10
	const b int16 = 20
	fmt.Printf("%v, %T\n", a+b, a+b)

}
