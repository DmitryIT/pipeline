package main

import "github.com/DmitryIT/pipeline/figures"

func GenerateFigure(figureType figures.FigureType) figures.Figure {
	switch figureType {
	case figures.TriangleType:
		return figures.NewTriangle()
	case figures.RectangleType:
		return figures.NewRectangle()
	case figures.CircleType:
		return figures.NewCircle()
	}
	return nil
}
