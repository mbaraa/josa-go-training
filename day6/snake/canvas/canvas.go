package canvas

import (
	"errors"
	"strings"
)

type Point struct {
	X       int
	Y       int
	content string
}

func NewPoint(x, y int, content string) Point {
	return Point{
		X:       x,
		Y:       y,
		content: content,
	}
}

func (p Point) Content() string {
	return p.content
}

type Canvas struct {
	width     int
	height    int
	emptyCell string
	cells     [][]Point
}

func NewCanvas(width, height int, emptyCell string) *Canvas {
	width += 1
	height += 1

	cells := make([][]Point, height)
	for y := 0; y < height; y++ {
		cells[y] = make([]Point, width)
		for x := 0; x < width; x++ {
			cells[y][x] = NewPoint(x, y, emptyCell)
		}
	}

	return &Canvas{
		width:     width,
		height:    height,
		emptyCell: emptyCell,
		cells:     cells,
	}
}

func (c *Canvas) String() string {
	builder := new(strings.Builder)
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			builder.WriteString(c.cells[y][x].Content())
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

func (c *Canvas) Size() Point {
	return NewPoint(c.width, c.height, "")
}

func (c *Canvas) MarkPoint(point Point) error {
	if !c.markable(point.X, point.Y) {
		return errors.New("point overflows the canvas")
	}

	c.cells[point.Y][point.X] = point
	return nil
}

func (c *Canvas) UnmarkPoint(point Point) error {
	if !c.markable(point.X, point.Y) {
		return errors.New("point overflows the canvas")
	}

	c.cells[point.Y][point.X] = NewPoint(point.X, point.Y, c.emptyCell)
	return nil
}

func (c *Canvas) Marked(x, y int) (Point, bool) {
	if !c.markable(x, y) {
		return Point{}, false
	}
	pt := c.cells[y][x]
	if pt.Content() == c.emptyCell {
		return Point{}, false
	}

	return pt, true
}

func (c *Canvas) markable(x, y int) bool {
	return x >= 0 &&
		x <= c.width &&
		y >= 0 &&
		y <= c.height
}
