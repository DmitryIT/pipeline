package pipeline

import (
	"context"
	"fmt"

	"github.com/DmitryIT/pipeline/figures"
	"github.com/DmitryIT/pipeline/metrics"
	"github.com/segmentio/ksuid"
)

func ProcessFigure(ctx context.Context, processType figures.ProcessingType, figure <-chan figures.Figure, resultChan chan<- float64) {
	id := ksuid.New()
	fmt.Printf("Processor %s started\n", id.String())
	metrics.CurrentNumberOfProcessors.Inc()
	defer metrics.CurrentNumberOfProcessors.Dec()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Processor %s stopped\n", id.String())
			return
		default:
			f := <-figure
			switch processType {
			case figures.Area:
				resultChan <- f.Area()
				metrics.FiguresProcessed.Inc()
			case figures.Perimeter:
				resultChan <- f.Perimeter()
				metrics.FiguresProcessed.Inc()
			}
			switch f.(type) {
			case *figures.Triangle:
				metrics.CurrentNumberOfTriangles.Dec()
			case *figures.Rectangle:
				metrics.CurrentNumberOfRectangles.Dec()
			case *figures.Circle:
				metrics.CurrentNumberOfCircles.Dec()
			}
			Delay()
		}
	}
}
