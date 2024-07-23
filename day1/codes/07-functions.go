package main

import (
	"fmt"
)

func main() {
	greet()
	greetName("Mansour")
	fmt.Println(cube(3))
	add := func(i, j int) int {
		return i + j
	}
	fmt.Println(calc(1, 2, add))
}

func greet() {
	fmt.Println("Hello World!")
}

func greetName(name string) {
	fmt.Println("Hello", name+"!")
}

func cube(i int) int {
	return i * i * i
}

func calc(x, y int, calculator func(i, j int) int) int {
	return calculator(x, y)
}
