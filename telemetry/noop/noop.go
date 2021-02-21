package telemetry

// NoopPusher implements MetricPusher interface with mock functions
type NoopPusher struct{}

func (nope *NoopPusher) Register(name string, desc string) {}

func (nope *NoopPusher) Configure() error {
	return nil
}

func (nope *NoopPusher) AddLabel(key string, value string) {}

func (nope *NoopPusher) Increment(name string) {}

func (nope *NoopPusher) Add(name string, delta float64) {}

func (nope *NoopPusher) SetToCurrentTime(name string) {}

func (nope *NoopPusher) Reset(name string) {}

func (nope *NoopPusher) Push() error {
	return nil
}
