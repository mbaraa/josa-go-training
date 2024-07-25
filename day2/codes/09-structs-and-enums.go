package main

import "fmt"

func main() {
	car := Car{}
	car.EngineType = "V9"
	car.Model = "2007"
	car.NumWheels = 4

	fmt.Printf("%s", car.String())
}

// public class Car {}

type EngineType string

const (
	V6  EngineType = "V6"
	V8  EngineType = "V8"
	V10 EngineType = "V10"
	V12 EngineType = "V12"
)

type Car struct {
	Model      string
	EngineType EngineType
	NumWheels  int
}

func (c *Car) bla() {
	fmt.Println(c)
}

func (c *Car) String() string {
	out := fmt.Sprintf("{\"model\": %s, \"engine_type\": %s, \"num_wheels\": %d}",
		c.Model, c.EngineType, c.NumWheels,
	)
	return out
}
