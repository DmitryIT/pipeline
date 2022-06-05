package figures

import (
	"math/rand"
	"time"
)

// define types of Figures
type FigureType string // "triangle", "rectangle", "circle"
const (
	TriangleType  FigureType = "triangle"
	RectangleType FigureType = "rectangle"
	CircleType    FigureType = "circle"
)

func RandomFugureType() FigureType {
	rand.Seed(time.Now().UnixNano())
	types := []FigureType{TriangleType, RectangleType, CircleType}
	return types[rand.Intn(len(types))]

}

type ProcessingType string

const (
	Area      ProcessingType = "area"
	Perimeter ProcessingType = "perimeter"
)

func RandomProcessingType() ProcessingType {
	rand.Seed(time.Now().UnixNano())
	types := []ProcessingType{Area, Perimeter}
	return types[rand.Intn(len(types))]
}

type Figure interface {
	Area() float64
	Perimeter() float64
}
