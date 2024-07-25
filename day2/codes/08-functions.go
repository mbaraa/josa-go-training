package main

import "fmt"

func main() {
	// &
	// *

	var i int = 10
	var ptrI *int = &i

	fmt.Println(i)
	fmt.Println(*ptrI)

	fmt.Println("---")

	y := 19
	fmt.Println("y", y)
	changeVal(y)
	fmt.Println("y changeVal", y)
	changeRef(&y)
	fmt.Println("y changeRef", y)
}

func changeVal(x int) {
	x += 10
}

func changeRef(x *int) {
	(*x) += 10
}
