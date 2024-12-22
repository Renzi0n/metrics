package main

import (
	"net/http"

	"github.com/Renzi0n/metrics/internal/handlers"
	memorystorage "github.com/Renzi0n/metrics/internal/memory_storage"
)

func main() {
	metricsStorage := memorystorage.NewMetricsStorage()
	metricsHandler := handlers.NewMetricsHandler(metricsStorage)
	mux := http.NewServeMux()
	mux.HandleFunc(`POST /update/{type}/{name}/{value}`, metricsHandler.UpdateMetrics)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
