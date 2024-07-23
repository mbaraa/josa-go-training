package main

import "fmt"

func main() {
	// int[] array = new int[x];
	var array [5]int
	array[0] = 1
	array[2] = 19
	array[4] = 18
	// array[10] = 12 // illegal

	for i := 0; i < len(array); i++ {
		fmt.Scan(&array[i])
	}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	fmt.Println("----")

	var slice []int
	fmt.Println("slice's length", len(slice))
	slice = append(slice, 12)
	slice = append(slice, 15)
	slice = append(slice, 17)
	slice = append(slice, -1)
	fmt.Println("slice's length", len(slice))
	fmt.Println("first element", slice[1])

	for index, value := range slice {
		fmt.Println("idx:", index, "val:", value)
	}
}
