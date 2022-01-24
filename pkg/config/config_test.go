package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	manager "github.ibm.com/chandergovind/mq2p/pkg/metric-manager"
)

func TestSimpleConfigMQPortion(t *testing.T) {
	assert := assert.New(t)

	configString := []byte(`
mq:
  server: dummy_server
  topic: dummy_topic
`)
	cfg, err := ParseConfig(configString)
	assert.Nil(err)
	assert.Equal("dummy_server", cfg.MqConfig.Server)
	assert.Equal("dummy_topic", cfg.MqConfig.Topic)
	assert.Empty(cfg.Metrics)

}

func TestSimpleConfigSingleMetricWithNoLabels(t *testing.T) {
	assert := assert.New(t)

	configString := []byte(`
mq:
  server: dummy_server
metrics:
  - mq_name: simple_metric
    prom_name: simple_metric_prom
    help: simple help string
    type: gauge
`)

	cfg, err := ParseConfig(configString)
	assert.Nil(err)
	assert.Equal(1, len(cfg.Metrics))

	metric := cfg.Metrics[0]
	assert.Equal("simple_metric", metric.MQName)
	assert.Equal("simple_metric_prom", metric.PromName)
	assert.Equal("simple help string", metric.Help)
	assert.Equal(manager.Gauge, metric.Type)
	assert.Empty(metric.Labels)
}

func TestSimpleConfigSingleMetric(t *testing.T) {
	assert := assert.New(t)

	configString := []byte(`
mq:
  server: dummy_server
metrics:
  - mq_name: simple_metric
    prom_name: simple_metric_prom
    help: simple help string
    type: gauge
    labels:
    - label1
    - label2
`)

	cfg, err := ParseConfig(configString)
	assert.Nil(err)
	assert.Equal(1, len(cfg.Metrics))

	metric := cfg.Metrics[0]
	assert.Equal(2, len(metric.Labels))
	assert.Equal("label1", metric.Labels[0])
	assert.Equal("label2", metric.Labels[1])
}

func TestSimpleConfigMultipleMetrics(t *testing.T) {
	assert := assert.New(t)

	configString := []byte(`
mq:
  server: dummy_server
metrics:
  - mq_name: simple_metric_1
    prom_name: simple_metric_prom_1
    help: simple help string
    type: gauge
    labels:
    - label1
    - label2
  - mq_name: simple_metric_2
    prom_name: simple_metric_prom_2
    help: simple help string
    type: gauge
    labels:
    - label3
    - label4
    - label5
`)

	cfg, err := ParseConfig(configString)
	assert.Nil(err)
	assert.Equal(2, len(cfg.Metrics))

	metric1 := cfg.Metrics[0]
	assert.Equal("simple_metric_1", metric1.MQName)
	assert.Equal(2, len(metric1.Labels))

	metric2 := cfg.Metrics[1]
	assert.Equal("simple_metric_2", metric2.MQName)
	assert.Equal(3, len(metric2.Labels))
}
