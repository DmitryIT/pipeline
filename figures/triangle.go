package figures

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

func (t *Triangle) Perimeter() float64 {
	// TODO: calculate real perimeter
	return 1
}
