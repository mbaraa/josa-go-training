package main

import (
	"fmt"
	"os"
	"snake/canvas"
	"snake/game"
	"snake/keyboard"
	"snake/utils"
)

func main() {
	termWidth, termHeight, err := utils.GetTermSize()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c := canvas.NewCanvas(termWidth-10, termHeight-5, " ")
	snake := game.NewSnake(8, "#", canvas.NewPoint(c.MaxPoint().X/2, c.MaxPoint().Y/2, "#"))
	gameThingy := game.NewGame(snake, c)

	kb := keyboard.New()
	kb.AddEventListener(keyboard.KeyArrowDown, func() {
		gameThingy.SetMoveDir(game.MoveDown)
	})
	kb.AddEventListener(keyboard.KeyArrowUp, func() {
		gameThingy.SetMoveDir(game.MoveUp)
	})
	kb.AddEventListener(keyboard.KeyArrowLeft, func() {
		gameThingy.SetMoveDir(game.MoveLeft)
	})
	kb.AddEventListener(keyboard.KeyArrowRight, func() {
		gameThingy.SetMoveDir(game.MoveRight)
	})

	go kb.Listen()
	gameThingy.LoopGame(func() {
		kb.Close()
	})
}
