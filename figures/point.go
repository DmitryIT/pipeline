package figures

import (
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
