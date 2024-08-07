package pgsqlmon

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"

	"gitlab.com/blog/ops/pkg/metrics"
)

// Copied from github.com/prometheus/client_golang/prometheus/collectors.
type dbStatsCollector struct {
	db *sql.DB

	maxOpenConnections *metrics.Desc
	openConnections    *metrics.Desc
	inUseConnections   *metrics.Desc
	idleConnections    *metrics.Desc
	waitCount          *metrics.Desc
	waitDuration       *metrics.Desc
	maxIdleClosed      *metrics.Desc
	maxIdleTimeClosed  *metrics.Desc
	maxLifetimeClosed  *metrics.Desc
}

// RegisterStatsCollector registers postgres stats collector.
func RegisterStatsCollector(db *sql.DB, dbName string) {
	collector := dbStatsCollector{
		db: db,
		maxOpenConnections: metrics.NewDesc(
			"sql_max_open_connections",
			"Maximum number of open connections to the database.",
			nil, metrics.Labels{"db_name": dbName},
		),
		openConnections: metrics.NewDesc(
			"sql_open_connections",
			"The number of established connections both in use and idle.",
			nil, metrics.Labels{"db_name": dbName},
		),
		inUseConnections: metrics.NewDesc(
			"sql_in_use_connections",
			"The number of connections currently in use.",
			nil, metrics.Labels{"db_name": dbName},
		),
		idleConnections: metrics.NewDesc(
			"sql_idle_connections",
			"The number of idle connections.",
			nil, metrics.Labels{"db_name": dbName},
		),
		waitCount: metrics.NewDesc(
			"sql_wait_count_total",
			"The total number of connections waited for.",
			nil, metrics.Labels{"db_name": dbName},
		),
		waitDuration: metrics.NewDesc(
			"sql_wait_duration_seconds_total",
			"The total time blocked waiting for a new connection.",
			nil, metrics.Labels{"db_name": dbName},
		),
		maxIdleClosed: metrics.NewDesc(
			"sql_max_idle_closed_total",
			"The total number of connections closed due to SetMaxIdleConns.",
			nil, metrics.Labels{"db_name": dbName},
		),
		maxIdleTimeClosed: metrics.NewDesc(
			"sql_max_idle_time_closed_total",
			"The total number of connections closed due to SetConnMaxIdleTime.",
			nil, metrics.Labels{"db_name": dbName},
		),
		maxLifetimeClosed: metrics.NewDesc(
			"sql_max_lifetime_closed_total",
			"The total number of connections closed due to SetConnMaxLifetime.",
			nil, metrics.Labels{"db_name": dbName},
		),
	}

	metrics.MustRegister(&collector)
}

// Describe implements Collector.
func (c *dbStatsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.maxOpenConnections
	ch <- c.openConnections
	ch <- c.inUseConnections
	ch <- c.idleConnections
	ch <- c.waitCount
	ch <- c.waitDuration
	ch <- c.maxIdleClosed
	ch <- c.maxLifetimeClosed
	ch <- c.maxIdleTimeClosed
}

// Collect implements Collector.
func (c *dbStatsCollector) Collect(ch chan<- prometheus.Metric) {
	stats := c.db.Stats()
	ch <- prometheus.MustNewConstMetric(c.maxOpenConnections, prometheus.GaugeValue, float64(stats.MaxOpenConnections))
	ch <- prometheus.MustNewConstMetric(c.openConnections, prometheus.GaugeValue, float64(stats.OpenConnections))
	ch <- prometheus.MustNewConstMetric(c.inUseConnections, prometheus.GaugeValue, float64(stats.InUse))
	ch <- prometheus.MustNewConstMetric(c.idleConnections, prometheus.GaugeValue, float64(stats.Idle))
	ch <- prometheus.MustNewConstMetric(c.waitCount, prometheus.CounterValue, float64(stats.WaitCount))
	ch <- prometheus.MustNewConstMetric(c.waitDuration, prometheus.CounterValue, stats.WaitDuration.Seconds())
	ch <- prometheus.MustNewConstMetric(c.maxIdleClosed, prometheus.CounterValue, float64(stats.MaxIdleClosed))
	ch <- prometheus.MustNewConstMetric(c.maxLifetimeClosed, prometheus.CounterValue, float64(stats.MaxLifetimeClosed))
	ch <- prometheus.MustNewConstMetric(c.maxIdleTimeClosed, prometheus.CounterValue, float64(stats.MaxIdleTimeClosed))
}
