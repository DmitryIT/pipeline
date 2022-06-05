package main

import "github.com/DmitryIT/pipeline/figures"

type ProcessingType string

const (
	Area      ProcessingType = "area"
	Perimeter ProcessingType = "perimeter"
)

func ProcessFigure(processType ProcessingType, figure <-chan figures.Figure, resultChan chan<- float64) {
	switch processType {
	case Area:
		resultChan <- (<-figure).Area()
	case Perimeter:
		resultChan <- (<-figure).Perimeter()
	}
}
