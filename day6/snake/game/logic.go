package game

import (
	"errors"
	"fmt"
	"math/rand"
	"snake/canvas"
	"snake/utils"
	"time"
)

type MoveDirection uint

const (
	MoveUp MoveDirection = iota
	MoveDown
	MoveLeft
	MoveRight
)

type GameLogic struct {
	score         int
	applePosition canvas.Point
	canvass       *canvas.Canvas
	snake         *Snake
	moveDir       MoveDirection
}

func NewGame(snake *Snake, c *canvas.Canvas) *GameLogic {
	game := &GameLogic{
		score:   0,
		snake:   snake,
		canvass: c,
		moveDir: MoveRight,
	}
	game.changeApplePosition()

	return game
}

func (g *GameLogic) drawSnake() error {
	for n := g.snake.Nodes().Front(); n.Next() != nil; n = n.Next() {
		snakePoint := n.Value.(*canvas.Point)
		err := g.canvass.MarkPoint(*snakePoint)
		if err != nil {
			return errors.New("snake hit a wall")
		}
		err = g.canvass.UnmarkPoint(g.snake.Tail())
		if err != nil {
			return errors.New("snake hit a wall")
		}
	}
	return nil
}

func (g *GameLogic) changeApplePosition() {
	randomizer := rand.New(rand.NewSource(time.Now().UnixMicro()))
	g.applePosition = canvas.NewPoint(
		randomizer.Intn(g.canvass.MaxPoint().X),
		randomizer.Intn(g.canvass.MaxPoint().Y),
		"A",
	)
	g.canvass.MarkPoint(g.applePosition)
}

func (g *GameLogic) SetMoveDir(dir MoveDirection) {
	g.moveDir = dir
}

func (g *GameLogic) LoopGame(exitCallBack func()) {
	timer := time.NewTicker(time.Second / 8)
	for range timer.C {
		switch g.moveDir {
		case MoveUp:
			g.snake.MoveUp()
		case MoveDown:
			g.snake.MoveDown()
		case MoveLeft:
			g.snake.MoveLeft()
		case MoveRight:
			g.snake.MoveRight()
		}

		var err error
		snakeHead := g.snake.Head()

		if content, marked := g.canvass.Marked(snakeHead); marked {
			switch content {
			case g.applePosition.Content():
				g.score++
				g.changeApplePosition()
			case g.snake.Content():
				g.Exit()
				exitCallBack()
				return
			}
		}
		err = g.drawSnake()
		if err != nil {
			g.Exit()
			exitCallBack()
			return
		}

		utils.ClearScreen()
		fmt.Println(g.canvass.String())
		fmt.Println("Score", g.score)
	}
}

func (g *GameLogic) Exit() {
	utils.ClearScreen()
	fmt.Println("GAME OVER")
	fmt.Println("Your final score was ", g.score)
}
