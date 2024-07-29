package game

import (
	"fmt"
	"math/rand"
	"snake/canvas"
	"snake/cfmt"
	"snake/utils"
	"time"
)

type MoveDirection int

const (
	MoveUp MoveDirection = iota
	MoveDown
	MoveLeft
	MoveRight
)

type GameLogic struct {
	score   int
	apple   canvas.Point
	can     *canvas.Canvas
	snake   *Snake
	moveDir MoveDirection
}

func NewGame(snake *Snake, can *canvas.Canvas) *GameLogic {
	game := &GameLogic{
		score:   0,
		can:     can,
		snake:   snake,
		moveDir: MoveRight,
	}

	game.updateApplePosition()

	return game
}

func (g *GameLogic) updateApplePosition() {
	random := rand.New(rand.NewSource(time.Now().UnixMicro()))

	var appleX, appleY int

	nodes := g.snake.Nodes()
	for node := nodes.Front(); node.Next() != nil; node = node.Next() {
		pt := node.Value.(canvas.Point)

		appleX = random.Intn(g.can.Size().X)
		appleY = random.Intn(g.can.Size().Y)

		if pt.X == appleX && pt.Y == appleY {
			continue
		} else {
			break
		}
	}

	g.apple = canvas.NewPoint(appleX, appleY, cfmt.Red().Bold().Sprint("O"))
	g.can.MarkPoint(g.apple)
}

func (g *GameLogic) drawSnake() error {
	nodes := g.snake.Nodes()
	for node := nodes.Front(); node.Next() != nil; node = node.Next() {
		pt := node.Value.(canvas.Point)
		err := g.can.MarkPoint(pt)
		if err != nil {
			return err
		}
		tail := g.snake.Tail()
		err = g.can.UnmarkPoint(tail)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *GameLogic) LoopGame(exitCallback func()) {
	ticker := time.NewTicker(time.Second / 8)
	for range ticker.C {
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

		snakeHead := g.snake.Head()
		point, marked := g.can.Marked(snakeHead.X, snakeHead.Y)
		if marked {
			switch point.Content() {
			case g.apple.Content():
				g.score += 1
				g.snake.AddNode()
				g.updateApplePosition()
			case g.snake.CellContent():
				g.Exit()
				exitCallback()
				return
			}
		}

		err := g.drawSnake()
		if err != nil {
			g.Exit()
			exitCallback()
			return
		}

		utils.ClearScreen()
		fmt.Println(g.can.String())
		fmt.Println("Score: ", g.score)
	}
}

func (g *GameLogic) Exit() {
	utils.ClearScreen()
	cfmt.Red().Println("GAME OVER")
	cfmt.Blue().Println("Your final score was ", g.score)
}

func (g *GameLogic) SetMoveDirection(dir MoveDirection) {
	g.moveDir = dir
}
