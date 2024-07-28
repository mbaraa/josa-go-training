package canvas

import "strings"

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
