package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/DmitryIT/pipeline/pipeline"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var numOfGenerators, numOfProcessors uint64
	// take number of generators from enviroment variable
	if numOfGeneratorsStr, found := os.LookupEnv("NUM_OF_GENERATORS"); found {
		numOfGenerators, _ = strconv.ParseUint(numOfGeneratorsStr, 10, 64)
	} else {
		numOfGenerators = 1
	}
	// take number of processors from enviroment variable
	if numOfProcessorsStr, found := os.LookupEnv("NUM_OF_PROCESSORS"); found {
		numOfProcessors, _ = strconv.ParseUint(numOfProcessorsStr, 10, 64)
	} else {
		numOfProcessors = 1
	}

	pipeline.StartPipeline(uint(numOfGenerators), uint(numOfProcessors))

	fmt.Printf("Starting server on port 8080\n")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
