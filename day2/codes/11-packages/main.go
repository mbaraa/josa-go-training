package main

import (
	"fmt"
	"pkgs/utils"
	"time"

	"github.com/google/uuid"
)

func main() {
	fmt.Println(uuid.NewString())

	fmt.Println("Hello world")

	// utils.hello() // CAN'T ACCESS A PRIVATE (UNEXPORTED) FUNCTION

	time.Sleep(time.Second * 2)
	utils.ClearScreen()
}
