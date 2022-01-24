package manager

import (
	"errors"

	"gopkg.in/yaml.v3"
)

type MetricSpec struct {
	Metrics []MetricConfig
}

type MetricType int64

const (
	Gauge MetricType = iota
	Counter
)

type MetricConfig struct {
	MQName   string `yaml:"mq_name"`
	PromName string `yaml:"prom_name"`
	Help     string
	Type     MetricType
	Labels   []string
}

type MQPayload []MetricPayload

type MetricPayload struct {
	Name   string
	Value  float64
	Labels map[string]string
}

func (mt *MetricType) UnmarshalYAML(node *yaml.Node) error {
	switch node.Value {
	case "gauge":
		*mt = Gauge
	case "counter":
		*mt = Counter
	default:
		return errors.New("Unsupported Metric Type found. Cannot unmarshal.")
	}

	return nil
}
