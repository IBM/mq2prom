package manager

type Manager interface {
	// Initialize the Metric Manager with a list of metrics to track
	Init(MetricSpec)
	Update(MQPayload)
}
