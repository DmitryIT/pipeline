package figures

// define types of Figures
type FigureType string // "triangle", "rectangle", "circle"
const (
	TriangleType  FigureType = "triangle"
	RectangleType FigureType = "rectangle"
	CircleType    FigureType = "circle"
)

type Figure interface {
	Area() float64
	Perimeter() float64
}
