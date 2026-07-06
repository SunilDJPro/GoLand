package main

import "fmt"

// Define a struct
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle implements Shape interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// Circle implements Shape interface
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}

	fmt.Print("Rectangle ")
	printArea(rect)
	fmt.Print("Circle ")
	printArea(circ)
}
