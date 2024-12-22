package handlers

import (
	"net/http"

	memorystorage "github.com/Renzi0n/metrics/internal/memory_storage"
	"github.com/Renzi0n/metrics/internal/models"
)

type MetricsHandler struct {
	MetricsStorage *memorystorage.MetricsStorage
}

func NewMetricsHandler(MetricsStorage *memorystorage.MetricsStorage) *MetricsHandler {
	return &MetricsHandler{
		MetricsStorage: MetricsStorage,
	}
}

func (m *MetricsHandler) UpdateMetrics(res http.ResponseWriter, req *http.Request) {
	metricType := models.MetricType(req.PathValue("type"))
	metricName := req.PathValue("name")
	metricValue := req.PathValue("value")

	if metricType != models.Gauge && metricType != models.Counter {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	err := m.MetricsStorage.UpdateMetrics(metricType, metricName, metricValue)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
}
