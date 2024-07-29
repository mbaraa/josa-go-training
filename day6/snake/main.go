package main

import (
	"fmt"
	"os"
	"snake/canvas"
	"snake/cfmt"
	"snake/game"
	"snake/keyboard"
	"snake/utils"
)

func main() {
	w, h, err := utils.GetTerminalSize()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	snake := game.NewSnake(10, canvas.NewPoint(20, 12, cfmt.Green().Bold().Sprint("S")))
	can := canvas.NewCanvas(w-1, h-5, " ")

	gameLogic := game.NewGame(snake, can)

	kb := keyboard.New()
	kb.AddEventListener(keyboard.KeyArrowDown, func() {
		gameLogic.SetMoveDirection(game.MoveDown)
	})
	kb.AddEventListener(keyboard.KeyArrowUp, func() {
		gameLogic.SetMoveDirection(game.MoveUp)
	})
	kb.AddEventListener(keyboard.KeyArrowLeft, func() {
		gameLogic.SetMoveDirection(game.MoveLeft)
	})
	kb.AddEventListener(keyboard.KeyArrowRight, func() {
		gameLogic.SetMoveDirection(game.MoveRight)
	})

	go kb.Listen()
	gameLogic.LoopGame(func() {
		kb.Close()
	})
}
