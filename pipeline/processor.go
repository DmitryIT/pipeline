package pipeline

import (
	"context"
	"fmt"

	"github.com/DmitryIT/pipeline/figures"
	"github.com/segmentio/ksuid"
)

func ProcessFigure(ctx context.Context, processType figures.ProcessingType, figure <-chan figures.Figure, resultChan chan<- float64) {
	id := ksuid.New()
	fmt.Printf("Processor %s started\n", id.String())

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Processor %s stopped\n", id.String())
			return
		default:
			switch processType {
			case figures.Area:
				resultChan <- (<-figure).Area()
			case figures.Perimeter:
				resultChan <- (<-figure).Perimeter()
			}
			Delay()
		}
	}
}
