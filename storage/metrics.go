package storage

import (
    "github.com/prometheus/client_golang/prometheus"
)

func NewPrometheusRegistry() *prometheus.Registry {
    registry := prometheus.NewRegistry()

	jobCountGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "queued_job_count",
		Help: "The current number of jobs in queue",
	})
	registry.MustRegister(jobCountGauge)

	jobCountGauge.Set(5)

	return registry
}