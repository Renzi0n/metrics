package models

type MetricType string

const (
	Gauge   MetricType = "gauge"
	Counter MetricType = "counter"
)

type GaugeMetricValue float64
type CounterMetricValue int64
