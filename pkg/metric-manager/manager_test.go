package manager

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

func TestManagerSingleGauge(t *testing.T) {
	assert := assert.New(t)

	sm := NewSimpleManager()

	// Setup a single metric - Gauge
	sm.Init(MetricSpec{
		Metrics: []MetricConfig{
			{MQName: "simple_metric", PromName: "simple_metric_prom", Help: "dummy metric value", Type: Gauge, Labels: nil},
		},
	})

	// Post an update
	sm.Update([]MetricPayload{
		{Name: "simple_metric", Value: 10},
	})

	// Check if metric reflects the update
	assert.Equal(float64(10), testutil.ToFloat64(sm.metrics["simple_metric"]))

	// Change gauge value again
	sm.Update([]MetricPayload{
		{Name: "simple_metric", Value: 5},
	})

	// Check new value
	assert.Equal(float64(5), testutil.ToFloat64(sm.metrics["simple_metric"]))
}

func TestManagerTwoMetrics(t *testing.T) {
	assert := assert.New(t)

	sm := NewSimpleManager()

	// Setup a 2 metrics - a Gauge and a Counter
	sm.Init(MetricSpec{
		Metrics: []MetricConfig{
			{MQName: "simple_gauge", PromName: "simple_gauge_prom", Help: "simple gauge metric", Type: Gauge, Labels: nil},
			{MQName: "simple_counter", PromName: "simple_counter_prom", Help: "simple counter metric", Type: Counter, Labels: nil},
		},
	})

	// Post an update to both metrics in one shot
	sm.Update([]MetricPayload{
		{Name: "simple_gauge", Value: 10},
		{Name: "simple_counter", Value: 5},
	})

	// Check if metric reflects the update
	assert.Equal(float64(10), testutil.ToFloat64(sm.metrics["simple_gauge"]))
	assert.Equal(float64(5), testutil.ToFloat64(sm.metrics["simple_counter"]))

	// Change counter value alone
	sm.Update(MQPayload{
		{Name: "simple_counter", Value: 5},
	})

	// Gauge should not have changed
	assert.Equal(float64(10), testutil.ToFloat64(sm.metrics["simple_gauge"]))
	// Counter should have
	assert.Equal(float64(10), testutil.ToFloat64(sm.metrics["simple_counter"]))
}

func TestManagerSingleLabel(t *testing.T) {
	assert := assert.New(t)

	sm := NewSimpleManager()

	// Setup a single metric - a counter - with a single label
	sm.Init(MetricSpec{
		Metrics: []MetricConfig{
			{MQName: "label_metric_1", PromName: "label_metric_1_prom", Help: "metric value with labels", Type: Counter, Labels: []string{"label1"}},
		},
	})

	// Post an update
	sm.Update(MQPayload{
		{Name: "label_metric_1", Value: 10, Labels: map[string]string{"label1": "value1"}},
	})

	// Check if metric reflects the update
	assert.Equal(float64(10), testutil.ToFloat64(sm.metrics["label_metric_1"].(*prometheus.CounterVec).With(map[string]string{"label1": "value1"})))

	// Post an update
	sm.Update(MQPayload{
		{Name: "label_metric_1", Value: 10, Labels: map[string]string{"label1": "value2"}},
	})

	// check if both the counters have right values now
	assert.Equal(float64(10), testutil.ToFloat64(sm.metrics["label_metric_1"].(*prometheus.CounterVec).With(map[string]string{"label1": "value1"})))
	assert.Equal(float64(10), testutil.ToFloat64(sm.metrics["label_metric_1"].(*prometheus.CounterVec).With(map[string]string{"label1": "value2"})))

	// Post a double update
	sm.Update(MQPayload{
		{Name: "label_metric_1", Value: 10, Labels: map[string]string{"label1": "value1"}},
		{Name: "label_metric_1", Value: 10, Labels: map[string]string{"label1": "value2"}},
	})

	// check if both counters are updated
	assert.Equal(float64(20), testutil.ToFloat64(sm.metrics["label_metric_1"].(*prometheus.CounterVec).With(map[string]string{"label1": "value1"})))
	assert.Equal(float64(20), testutil.ToFloat64(sm.metrics["label_metric_1"].(*prometheus.CounterVec).With(map[string]string{"label1": "value2"})))
}

func TestManagerMultipleLabels(t *testing.T) {
	assert := assert.New(t)

	sm := NewSimpleManager()

	// Setup a single metric - a gauge - with a 3 labels
	sm.Init(MetricSpec{
		Metrics: []MetricConfig{
			{MQName: "label_metric_2", PromName: "label_metric_2_prom", Help: "metric value with multiple labels", Type: Gauge, Labels: []string{"label1", "label2", "label3"}},
		},
	})

	// Post a bunch of updates
	sm.Update(MQPayload{
		{Name: "label_metric_2", Value: 10, Labels: map[string]string{"label1": "value11", "label2": "value21", "label3": "value31"}},
		{Name: "label_metric_2", Value: 20, Labels: map[string]string{"label1": "value12", "label2": "value22", "label3": "value32"}},
	})

	assert.Equal(float64(10), testutil.ToFloat64(sm.metrics["label_metric_2"].(*prometheus.GaugeVec).With(map[string]string{"label1": "value11", "label2": "value21", "label3": "value31"})))
	assert.Equal(float64(20), testutil.ToFloat64(sm.metrics["label_metric_2"].(*prometheus.GaugeVec).With(map[string]string{"label1": "value12", "label2": "value22", "label3": "value32"})))
}
