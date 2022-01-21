package manager

type MetricSpec struct {
	Metrics []MetricConfig
}

type MetricType int64

const (
	Gauge MetricType = iota
	Counter
)

type MetricConfig struct {
	MQName   string
	PromName string
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
