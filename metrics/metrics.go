package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	FiguresProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "pipeline_processed_figures_total",
		Help: "The total number of processed figures",
	})
	CurrentNumberOfTriangles = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "pipeline_current_number_of_triangles",
		Help: "The current number of triangles",
	})
	CurrentNumberOfRectangles = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "pipeline_current_number_of_rectangles",
		Help: "The current number of rectangles",
	})
	CurrentNumberOfCircles = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "pipeline_current_number_of_circles",
		Help: "The current number of circles",
	})
)

func recordMetrics() {
	go func() {
		for {
			FiguresProcessed.Inc()
			time.Sleep(time.Second)
		}
	}()
}
