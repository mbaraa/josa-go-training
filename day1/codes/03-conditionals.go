package main

import "fmt"

func main() {
	shallPass := 10 > 11

	if shallPass {
		fmt.Println("You shall pass")
	} else {
		fmt.Println("You shall not pass")
	}

	var age int
	fmt.Print("Enter a number: ")
	fmt.Scan(&age)

	switch age {
	case 4:
		fmt.Println("can go to school")
	case 16, 18:
		fmt.Println("can drive a car")
	case 21:
		fmt.Println("can vote")
	case 60:
		fmt.Println("can retire")
	default:
		fmt.Println("invalid age")
	}
}
