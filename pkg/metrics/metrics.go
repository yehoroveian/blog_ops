package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// DefaultNamespace is a common namespace for all metrics specified in Default services.
const DefaultNamespace = "default"

// DefaultMetricsEndpoint is a default export path for prometheus crawler.
const DefaultMetricsEndpoint = "/metrics"

var (
	registerer = prometheus.DefaultRegisterer
	gatherer   = prometheus.DefaultGatherer
)

var (
	// DefaultSummaryObjectives are default percentiles for 50%, 90% and 99%.
	DefaultSummaryObjectives = map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}

	// DefBuckets are the default Histogram buckets.
	DefBuckets = prometheus.DefBuckets
)

type (
	// Labels represents a collection of label name -> value mappings.
	Labels = prometheus.Labels

	// Desc is the descriptor used by every Prometheus Metric.
	Desc = prometheus.Desc

	// Collector is the interface implemented by anything that can be used by
	// Prometheus to collect metrics.
	Collector = prometheus.Collector
)

// Handler returns HTTP handler for serving metrics.
func Handler() http.Handler {
	return promhttp.Handler()
}

// Gatherer returns default gatherer interface from Prometheus.
func Gatherer() prometheus.Gatherer {
	return gatherer
}

// Registerer returns default registry from the Prometheus.
// It is the most common way to use the client.
func Registerer() prometheus.Registerer {
	return registerer
}

// MustRegister registers a new metric in the registry. If the metric fail to register it will panic.
func MustRegister(collectors ...Collector) {
	registerer.MustRegister(collectors...)
}

// Unregister metric from the registry.
func Unregister(collector Collector) bool {
	return registerer.Unregister(collector)
}

// NewCounter creates a new Counter with predefined namespace
func NewCounter(name, help string) prometheus.Counter {
	return prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: DefaultNamespace,
			Name:      name,
			Help:      help,
		},
	)
}

// NewGauge creates a new Gauge with predefined namespace
func NewGauge(name, help string) prometheus.Gauge {
	return prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: DefaultNamespace,
			Name:      name,
			Help:      help,
		},
	)
}

// NewHistogram creates a new Histogram with predefined namespace
func NewHistogram(name, help string, buckets []float64) prometheus.Histogram {
	return prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: DefaultNamespace,
			Name:      name,
			Help:      help,
			Buckets:   buckets,
		},
	)
}

// NewSummary creates a new Summary with predefined namespace.
func NewSummary(name, help string) prometheus.Summary {
	return prometheus.NewSummary(
		prometheus.SummaryOpts{
			Namespace:  DefaultNamespace,
			Name:       name,
			Help:       help,
			MaxAge:     time.Minute,
			Objectives: DefaultSummaryObjectives,
		},
	)
}

// NewCounterVec creates a new CounterVec with predefined namespace.
func NewCounterVec(name, help string, labelValues []string) *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: DefaultNamespace,
			Name:      name,
			Help:      help,
		},
		labelValues,
	)
}

// NewGaugeVec creates a new GaugeVec with predefined namespace.
func NewGaugeVec(name, help string, labelValues []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: DefaultNamespace,
			Name:      name,
			Help:      help,
		},
		labelValues,
	)
}

// NewHistogramVec creates a new HistogramVec with predefined namespace.
func NewHistogramVec(name, help string, buckets []float64, labelValues []string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: DefaultNamespace,
			Name:      name,
			Help:      help,
			Buckets:   buckets,
		},
		labelValues,
	)
}

// NewSummaryVec creates a new SummaryVec with predefined namespace
func NewSummaryVec(name, help string, labelValues []string) *prometheus.SummaryVec {
	return prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace:  DefaultNamespace,
			Name:       name,
			Help:       help,
			MaxAge:     time.Minute,
			Objectives: DefaultSummaryObjectives,
		},
		labelValues,
	)
}

// MustRegisterCounter creates and registers new Counter with predefined namespace
// Panics if metrics with same name already registered
func MustRegisterCounter(name, help string) prometheus.Counter {
	collector := NewCounter(name, help)
	MustRegister(collector)

	return collector
}

// MustRegisterGauge creates and registers a new Gauge with predefined namespace.
// Panics if metrics with the same name are already registered.
func MustRegisterGauge(name, help string) prometheus.Gauge {
	collector := NewGauge(name, help)
	MustRegister(collector)

	return collector
}

// MustRegisterHistogram creates and registers new Histogram with predefined namespace.
// Panics if metrics with the same name are already registered.
func MustRegisterHistogram(name, help string, buckets []float64) prometheus.Histogram {
	collector := NewHistogram(name, help, buckets)
	MustRegister(collector)

	return collector
}

// MustRegisterSummary creates and registers new Summary with predefined namespace.
// Panics if metrics with the same name are already registered.
func MustRegisterSummary(name, help string) prometheus.Summary {
	collector := NewSummary(name, help)
	MustRegister(collector)

	return collector
}

// MustRegisterCounterVec creates and registers new CounterVec with predefined namespace.
// Panics if metrics with the same name are already registered.
func MustRegisterCounterVec(name, help string, labelValues []string) *prometheus.CounterVec {
	collector := NewCounterVec(name, help, labelValues)
	MustRegister(collector)

	return collector
}

// MustRegisterGaugeVec creates and registers new GaugeVec with predefined namespace.
// Panics if metrics with the same name are already registered.
func MustRegisterGaugeVec(name, help string, labelValues []string) *prometheus.GaugeVec {
	collector := NewGaugeVec(name, help, labelValues)
	MustRegister(collector)

	return collector
}

// MustRegisterHistogramVec creates and registers new HistogramVec with predefined namespace.
// Panics if metrics with the same name are already registered.
func MustRegisterHistogramVec(name, help string, buckets []float64, labelValues []string) *prometheus.HistogramVec {
	collector := NewHistogramVec(name, help, buckets, labelValues)
	MustRegister(collector)

	return collector
}

// MustRegisterSummaryVec creates and registers new SummaryVec with predefined namespace.
// Panics if metrics with the same name are already registered.
func MustRegisterSummaryVec(name, help string, labelValues []string) *prometheus.SummaryVec {
	collector := NewSummaryVec(name, help, labelValues)
	MustRegister(collector)

	return collector
}

// NewDesc allocates and initializes a new Desc. Errors are recorded in the Desc
// and will be reported on registration time. variableLabels and constLabels can
// be nil if no such labels should be set. fqName must not be empty.
//
// variableLabels only contain the label names. Their label values are variable
// and therefore not part of the Desc. (They are managed within the Metric.)
//
// For constLabels, the label values are constant. Therefore, they are fully
// specified in the Desc. See the Collector example for a usage pattern.
func NewDesc(fqName, help string, variableLabels []string, constLabels Labels) *prometheus.Desc {
	name := DefaultNamespace + "_" + fqName

	return prometheus.NewDesc(name, help, variableLabels, constLabels)
}
