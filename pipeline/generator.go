package pipeline

import (
	"context"
	"fmt"

	"github.com/DmitryIT/pipeline/figures"
	"github.com/segmentio/ksuid"
)

func GenerateFigure(ctx context.Context, figureType figures.FigureType, figureChan chan<- figures.Figure) {
	id := ksuid.New()
	fmt.Printf("Generator %s started\n", id.String())

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Generator %s stopped\n", id.String())
			return
		default:
			switch figureType {
			case figures.TriangleType:
				figureChan <- figures.NewTriangle()
			case figures.RectangleType:
				figureChan <- figures.NewRectangle()
			case figures.CircleType:
				figureChan <- figures.NewCircle()
			}
			Delay()
		}
	}
}
