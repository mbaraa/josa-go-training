package main

import "fmt"

func main() {
	var v Vehicle
	v = &Car{}
	v.Start()
	v = &Scooter{}
	v.Start()
	// v = Bus{} // illegal, bus doesn't implement Vehicle
}

type Vehicle interface {
	Start() error
	Stop() error
	Accelerate() error
	Brake() error
}

type Bus struct{}

type Car struct{}

func (c *Car) Start() error {
	fmt.Println("car started")
	return nil
}

func (c *Car) Stop() error {
	fmt.Println("car stopped")
	return nil
}

func (c *Car) Accelerate() error {
	fmt.Println("car is accelerating")
	return nil
}

func (c *Car) Brake() error {
	fmt.Println("car is braking")
	return nil
}

type Scooter struct{}

func (s *Scooter) Start() error {
	fmt.Println("scooter is staring")
	return nil
}

func (s *Scooter) Stop() error {
	panic("not implemented") // TODO: Implement
}

func (s *Scooter) Accelerate() error {
	panic("not implemented") // TODO: Implement
}

func (s *Scooter) Brake() error {
	panic("not implemented") // TODO: Implement
}
