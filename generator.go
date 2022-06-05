package main

import "github.com/DmitryIT/pipeline/figures"

func GenerateFigure(figureType figures.FigureType, figureChan chan<- figures.Figure) {
	switch figureType {
	case figures.TriangleType:
		figureChan <- figures.NewTriangle()
	case figures.RectangleType:
		figureChan <- figures.NewRectangle()
	case figures.CircleType:
		figureChan <- figures.NewCircle()
	}
}
