package main

import (
	"math"
	"math/rand"
	"time"
)

type Point struct {
	X int
	Y int
}

// returns a random Point where X and Y are between -1 and 1
func NewRandomPoint() *Point {
	rand.Seed(time.Now().UnixNano())
	return &Point{
		X: rand.Intn(2) - 1,
		Y: rand.Intn(2) - 1,
	}
}

type Figure interface {
	Area() float64
}

type Triangle struct {
	A Point
	B Point
	C Point
}

// returns a new Triangle
func NewTriangle() *Triangle {
	return &Triangle{
		A: *NewRandomPoint(),
		B: *NewRandomPoint(),
		C: *NewRandomPoint(),
	}
}

func (t *Triangle) Area() float64 {
	// TODO: calculate real area
	return 1
}

type Rectangle struct {
	A Point
	B Point
	C Point
	D Point
}

func NewRectangle() *Rectangle {
	return &Rectangle{
		A: *NewRandomPoint(),
		B: *NewRandomPoint(),
		C: *NewRandomPoint(),
		D: *NewRandomPoint(),
	}
}

func (r *Rectangle) Area() float64 {
	// TODO: calculate real area
	return 1
}

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
