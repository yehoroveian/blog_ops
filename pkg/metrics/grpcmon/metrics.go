package grpcmon

import (
	"time"

	"gitlab.com/blog/ops/pkg/metrics"
)

var (
	// grpcServerStartedCounter counts the number of started gRPC requests.
	grpcServerStartedCounter = metrics.NewCounterVec(
		"grpc_server_started_total",
		"Total number of RPCs started on the server.",
		[]string{"grpc_service", "grpc_method"},
	)

	// grpcServerStartedCounter counts the number of handled gRPC requests.
	grpcServerHandledCounter = metrics.NewCounterVec(
		"grpc_server_handled_total",
		"Total number of RPCs completed on the server, regardless of success or failure.",
		[]string{"grpc_service", "grpc_method", "grpc_code"},
	)

	// grpcServerHandledHistogram ccollects information about handling time of gRPC requests.
	grpcServerHandledHistogram = metrics.NewHistogramVec(
		"grpc_server_handling_seconds",
		"Histogram of response latency (seconds) of gRPC that had been application-level handled by the server.",
		metrics.DefBuckets,
		[]string{"grpc_service", "grpc_method"},
	)
)

func init() {
	metrics.MustRegister(
		grpcServerStartedCounter,
		grpcServerHandledCounter,
		grpcServerHandledHistogram,
	)
}

func IncServerStartedTotal(service string, method string) {
	grpcServerStartedCounter.WithLabelValues(service, method).Inc()
}

func IncServerHandledTotal(service, method, code string) {
	grpcServerHandledCounter.WithLabelValues(service, method, code).Inc()
}

func ObserveServerHandlingSeconds(service, method string, latency time.Duration) {
	grpcServerHandledHistogram.WithLabelValues(service, method).Observe(latency.Seconds())
}
