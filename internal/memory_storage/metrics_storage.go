package memorystorage

import (
	"errors"
	"strconv"

	"github.com/Renzi0n/metrics/internal/models"
)

type metrics struct {
	Gauge   map[string]models.GaugeMetricValue
	Counter map[string]models.CounterMetricValue
}

type MetricsStorage struct {
	Metrics metrics
}

func NewMetricsStorage() *MetricsStorage {
	return &MetricsStorage{
		Metrics: metrics{
			Gauge:   map[string]models.GaugeMetricValue{},
			Counter: map[string]models.CounterMetricValue{},
		},
	}
}

func (m *MetricsStorage) UpdateMetrics(Type models.MetricType, Name string, Value string) error {
	switch Type {
	case models.Gauge:
		val, err := strconv.ParseFloat(Value, 64)
		if err != nil {
			return err
		}
		m.Metrics.Gauge[Name] = models.GaugeMetricValue(val)
	case models.Counter:
		val, err := strconv.Atoi(Value)
		if err != nil {
			return err
		}
		m.Metrics.Counter[Name] += models.CounterMetricValue(val)
	default:
		return errors.New("WRONG METRIC TYPE")
	}

	return nil
}
