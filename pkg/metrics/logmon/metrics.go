package logmon

import (
	"gitlab.com/blog/ops/pkg/metrics"
)

var (
	logMessagesCounter = metrics.NewCounterVec(
		"log_messages_total",
		"Total number of messages written by logger.",
		[]string{"level"},
	)
)

func init() {
	metrics.MustRegister(logMessagesCounter)
}

// IncLogMessagesTotal increments log level counter.
func IncLogMessagesTotal(level string) {
	logMessagesCounter.WithLabelValues(level).Inc()
}
