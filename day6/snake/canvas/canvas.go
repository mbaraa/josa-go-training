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
	cells := make([][]Point, height+1)
	for y := 0; y < height+1; y++ {
		cells[y] = make([]Point, width+1)
		for x := 0; x < width+1; x++ {
			cells[y][x] = NewPoint(x, y, emptyCell)
		}
	}

	return &Canvas{
		width:     width + 1,
		height:    height + 1,
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

func (c *Canvas) EmptyCell() string {
	return c.emptyCell
}

// Marked checks if the given coordinates are empty or not,
// and returns the actual cell's content.
func (c *Canvas) Marked(point Point) (string, bool) {
	if !c.markable(point) {
		return "", false
	}
	pointOnPlane := c.cells[point.Y][point.X]
	if pointOnPlane.Content() == c.emptyCell {
		return pointOnPlane.Content(), false
	}
	return pointOnPlane.Content(), true
}

// MarkPoint sets the given coordinates to the given string.
func (c *Canvas) MarkPoint(point Point) error {
	if !c.markable(point) {
		return errors.New("point overflows the plane")
	}
	c.cells[point.Y][point.X] = point

	return nil
}

// UnmarkPoint sets the given coordinates to the plane's empty cell.
func (c *Canvas) UnmarkPoint(point Point) error {
	if !c.markable(point) {
		return errors.New("point overflows the plane")
	}
	c.cells[point.Y][point.X] = NewPoint(point.X, point.Y, c.emptyCell)

	return nil
}

// MaxPoint returns the plane's size.
func (c *Canvas) MaxPoint() Point {
	return Point{c.width, c.height, c.emptyCell}
}

func (c *Canvas) markable(point Point) bool {
	return point.X <= c.width &&
		point.X >= 0 &&
		point.Y <= c.height &&
		point.Y >= 0
}
