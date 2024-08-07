package goruntimemon

import (
	"gitlab.com/blog/ops/pkg/metrics"
)

// Set metrics that are required for all go applications.
func init() {
	metrics.MustRegister(recoveredPanicsCounter)
}

var (
	// recoveredPanicsCounter counts the number of panic recoveries from different locations, where panics raised.
	recoveredPanicsCounter = metrics.NewCounterVec(
		"go_recovered_panics_total",
		"A metric for a bundle of counters for recovered panics, labeled by location, where panic was raised.",
		[]string{"location"},
	)
)

// IncRecoveredPanics increases recoveredPanicsCounter by one.
func IncRecoveredPanics(location string) {
	recoveredPanicsCounter.WithLabelValues(location).Inc()
}
