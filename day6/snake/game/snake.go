package game

import (
	"container/list"
	"snake/canvas"
)

type Snake struct {
	cellContent string
	// List<canvas.Point>
	nodes *list.List
}

func NewSnake(numNodes int, position canvas.Point) *Snake {
	nodes := list.New()

	lastX := position.X
	for i := 0; i < numNodes; i++ {
		nodes.PushBack(canvas.NewPoint(lastX, position.Y, position.Content()))
		lastX -= 1
	}

	return &Snake{
		cellContent: position.Content(),
		nodes:       nodes,
	}
}

func (s *Snake) AddNode() {
	tail := s.Tail()
	tail.X -= 1
	s.nodes.PushBack(tail)
}

func (s *Snake) Nodes() list.List {
	return *s.nodes
}

func (s *Snake) MoveUp() {
	head := s.Head()
	s.move(head.X, head.Y-1)
}

func (s *Snake) MoveDown() {
	head := s.Head()
	s.move(head.X, head.Y+1)
}

func (s *Snake) MoveLeft() {
	head := s.Head()
	s.move(head.X-1, head.Y)
}

func (s *Snake) MoveRight() {
	head := s.Head()
	s.move(head.X+1, head.Y)
}

func (s *Snake) move(x, y int) {
	newHead := canvas.NewPoint(x, y, s.cellContent)
	tail := s.nodes.Back()

	s.nodes.PushFront(newHead)
	s.nodes.Remove(tail)
}

func (s *Snake) Head() canvas.Point {
	return s.nodes.Front().Value.(canvas.Point)
}

func (s *Snake) Tail() canvas.Point {
	return s.nodes.Back().Value.(canvas.Point)
}

func (s *Snake) CellContent() string {
	return s.cellContent
}
