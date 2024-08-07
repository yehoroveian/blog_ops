package tracemon

import (
	"github.com/prometheus/client_golang/prometheus"

	"gitlab.com/blog/ops/pkg/metrics"
)

var (
	tracerInitSuccess = metrics.NewGauge(
		"init_tracer_success",
		"Shows if error tracer init failed",
	)
)

func init() {
	prometheus.MustRegister(tracerInitSuccess)
}

// SetTracerInitSuccess explicitly sets value to 1 so that service
// is sure that tracer is initialized successfully (moved: 0 -> 1).
func SetTracerInitSuccess() {
	tracerInitSuccess.Set(1)
}
