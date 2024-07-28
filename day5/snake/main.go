package main

import (
	"fmt"
	"snake/canvas"
	"snake/keyboard"
)

func main() {
	can := canvas.NewCanvas(50, 25, "h")
	fmt.Println(can.String())

	return
	kb := keyboard.New()

	kb.AddEventListener(keyboard.KeyArrowDown, func() {
		fmt.Println("down")
	})
	kb.AddEventListener(keyboard.KeyArrowUp, func() {
		fmt.Println("up")
	})
	kb.AddEventListener(keyboard.KeyArrowLeft, func() {
		fmt.Println("left")
	})
	kb.AddEventListener(keyboard.KeyArrowRight, func() {
		fmt.Println("right")
	})

	go kb.Listen()
}
