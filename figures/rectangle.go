package figures

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

func (r *Rectangle) Perimeter() float64 {
	// TODO: calculate real perimeter
	return 1
}
