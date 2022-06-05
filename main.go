package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/DmitryIT/pipeline/figures"
)

const (
	NUM_OF_GENERATORS = 2
	NUM_OF_PROCESSORS = 0
)

func main() {
	figureChan := make(chan figures.Figure)
	resultChan := make(chan float64)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	// start up generators
	for i := 0; i < NUM_OF_GENERATORS; i++ {
		// choose random figure type
		go GenerateFigure(ctx, figures.RandomFugureType(), figureChan)
	}
	// start up processors
	for i := 0; i < NUM_OF_PROCESSORS; i++ {
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
				fmt.Printf("Result: %f\n", <-resultChan)
			}
		}
	}()
	<-ctx.Done()
}

// wait random time from 0 to 5 seconds
func Delay() {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
}
