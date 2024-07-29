package game

import (
	"container/list"
	"snake/canvas"
)

type Snake struct {
	position          canvas.Point
	singleCellContent string
	nodes             *list.List
}

func NewSnake(numNodes uint, content string, position canvas.Point) *Snake {
	nodes := new(list.List).Init()
	lastX := position.X
	for i := 0; i < int(numNodes); i++ {
		pt := canvas.NewPoint(lastX, position.Y, content)
		nodes.PushBack(&pt)
		lastX--
	}

	return &Snake{
		position:          position,
		singleCellContent: content,
		nodes:             nodes,
	}
}

func (s *Snake) Content() string {
	return s.singleCellContent
}

func (s *Snake) Nodes() *list.List {
	return s.nodes
}

func (s *Snake) Head() canvas.Point {
	return *s.nodes.Front().Value.(*canvas.Point)
}

func (s *Snake) Tail() canvas.Point {
	return *s.nodes.Back().Value.(*canvas.Point)
}

func (s *Snake) MoveUp() {
	head := s.nodes.Front().Value.(*canvas.Point)
	pt := canvas.NewPoint(head.X, head.Y-1, s.singleCellContent)
	s.move(&pt)
}

func (s *Snake) MoveDown() {
	head := s.nodes.Front().Value.(*canvas.Point)
	pt := canvas.NewPoint(head.X, head.Y+1, s.singleCellContent)
	s.move(&pt)
}

func (s *Snake) MoveLeft() {
	head := s.nodes.Front().Value.(*canvas.Point)
	pt := canvas.NewPoint(head.X-1, head.Y, s.singleCellContent)
	s.move(&pt)
}

func (s *Snake) MoveRight() {
	head := s.nodes.Front().Value.(*canvas.Point)
	pt := canvas.NewPoint(head.X+1, head.Y, s.singleCellContent)
	s.move(&pt)
}

func (s *Snake) move(point *canvas.Point) {
	s.nodes.PushFront(point)
	tail := s.nodes.Back()
	s.nodes.Remove(tail)
}
