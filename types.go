package main

import "math"

type Point struct {
	X int
	Y int
}

type Figure interface {
	Area() float64
}

type Triangle struct {
	A Point
	B Point
	C Point
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

func (r *Rectangle) Area() float64 {
	// TODO: calculate real area
	return 1
}

type Circle struct {
	Center Point
	Radius int
}

func (c *Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}
