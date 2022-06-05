package figures

import (
	"math"
	"math/rand"
)

type Circle struct {
	Center Point
	Radius int
}

func NewCircle() *Circle {
	return &Circle{
		Center: *NewRandomPoint(),
		Radius: rand.Intn(100),
	}
}

func (c *Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}

func (c *Circle) Perimeter() float64 {
	return math.Pi * float64(c.Radius) * 2
}
