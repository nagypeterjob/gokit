package telemetry

// MetricPusher provides an interface for Push based telemetry collection
type MetricPusher interface {
	// Perform any vendor specific setup.
	// If there is none, return nil
	Configure() error

	// Register gauge
	Register(name string, desc string)

	// Push gauges to storage backend
	Push() error

	// Add metric label
	AddLabel(key string, value string)

	// Increment gauge by one
	Increment(name string)

	// Increment gauge by arbitrary values
	Add(string, float64)

	// Set gauge to current unix time
	SetToCurrentTime(string)

	// Reset gauge
	Reset(string)
}
