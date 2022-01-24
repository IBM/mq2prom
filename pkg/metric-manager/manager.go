package manager

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type Manager interface {
	// Initialize the Metric Manager with a list of metrics to track
	Init(MetricSpec)
	// Update a set of metrics in a single message payload
	Update(MQPayload)
}

// This can be simplified into one array of struct that contains both the Type and the Collector
type SimpleManager struct {
	metrics     map[string]prometheus.Collector
	metricTypes map[string]MetricType
}

func NewSimpleManager() *SimpleManager {
	return &SimpleManager{
		metrics:     make(map[string]prometheus.Collector),
		metricTypes: make(map[string]MetricType),
	}
}

func (m *SimpleManager) Init(metricSpec MetricSpec) {
	for _, mc := range metricSpec.Metrics {
		var metricCollector prometheus.Collector

		switch mc.Type {
		case Gauge:
			metricCollector = prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: mc.PromName,
				Help: mc.Help,
			}, mc.Labels)
		case Counter:
			metricCollector = prometheus.NewCounterVec(prometheus.CounterOpts{
				Name: mc.PromName,
				Help: mc.Help,
			}, mc.Labels)
		}

		// register this metric with Prometheus
		prometheus.MustRegister(metricCollector)

		// store this metric collector for future use
		m.metrics[mc.MQName] = metricCollector
		m.metricTypes[mc.MQName] = mc.Type
	}
}

func (m *SimpleManager) Update(mqpayload MQPayload) {
	for _, payload := range mqpayload {
		collector := m.metrics[payload.Name]

		// Collector not found
		if collector == nil {
			fmt.Printf("Unrecognized metric found: %v. Skipping.", payload.Name)
			continue
		}

		switch m.metricTypes[payload.Name] {
		case Gauge:
			collector.(*prometheus.GaugeVec).With(payload.Labels).Set(payload.Value)
		case Counter:
			collector.(*prometheus.CounterVec).With(payload.Labels).Add(payload.Value)
		}
	}
}
