package utils

import "fmt"

// exported function

func ClearScreen() {
	fmt.Println("\033[H\033[2J")
}

// unexported function

func hello() {
}
