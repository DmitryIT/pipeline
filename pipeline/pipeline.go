package pipeline

import (
	"context"
	"math/rand"
	"time"

	"github.com/DmitryIT/pipeline/figures"
)

func StartPipeline(numOfGenerators uint, numOfProcessors uint) {

	figureChan := make(chan figures.Figure, numOfProcessors)
	resultChan := make(chan float64, numOfProcessors)
	ctx := context.Background()

	// start up generators
	for i := 0; i < int(numOfGenerators); i++ {
		// choose random figure type
		go GenerateFigure(ctx, figures.RandomFugureType(), figureChan)
	}
	// start up processors
	for i := 0; i < int(numOfProcessors); i++ {
		go ProcessFigure(ctx, figures.RandomProcessingType(), figureChan, resultChan)
	}

	// read from resultChan and print results
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(resultChan)
				return
			default:
				// fmt.Printf("Result: %f\n", <-resultChan)
				<-resultChan
			}
		}
	}()
}

// wait random time from 0 to 5 seconds
func Delay() {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
}
