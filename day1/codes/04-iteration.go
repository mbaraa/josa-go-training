package main

import "fmt"

func main() {
	// for
	// while
	// do while
	// foreach

	// for

	// for <init stmt>; <stop condition>; <update stmt> {
	// .. some code
	// }

	for i := 1; i <= 10; i++ {
		fmt.Println("Hello", i)
	}

	keepGoing := true
	i := 1
	for keepGoing {
		if i > 10 {
			keepGoing = false
		}
		fmt.Println("Hellloo")
		i++
	}

}
