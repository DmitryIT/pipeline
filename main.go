package main

import (
	"fmt"
	"net/http"

	"github.com/DmitryIT/pipeline/pipeline"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	pipeline.StartPipeline()

	fmt.Printf("Starting server on port 8080\n")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
