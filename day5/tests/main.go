package main

import "fmt"

func main() {
	m1 := [2][3]int{}
	// 1 row
	// 2 col
	// 2 rows
	// 3 cols
	m1[1][2] = 9

	fmt.Println("m1")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(m1[i][j], " ")
		}
		fmt.Println("")
	}

	m2 := [6]int{}
	// row*cols+col
	m2[1*3+2] = 10

	fmt.Println("m2")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(m2[i*3+j], " ")
		}
		fmt.Println("")
	}
}
