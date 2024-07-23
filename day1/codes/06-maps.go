package main

import "fmt"

func main() {
	// JAVA DON'T UNCOMMENT!!
	// HashMap shapes = new HashMap<String, Integer>()

	// sidesOfShapes := make(map[string]int)
	var sidesOfShapes map[string]int
	sidesOfShapes = make(map[string]int)

	sidesOfShapes["triangle"] = 3
	sidesOfShapes["square"] = 4
	sidesOfShapes["pentagon"] = 5
	sidesOfShapes["hexagon"] = 6
	sidesOfShapes["optagon"] = 7

	for index, value := range sidesOfShapes {
		fmt.Println("idx:", index, "val:", value)
	}
}
