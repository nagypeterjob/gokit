package telemetry

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

// PrometheusPusher implements MetricPusher interface
type PrometheusPusher struct {
	app       string
	metrics   map[string]*prometheus.GaugeVec
	labels    map[string]string
	labelKeys []string
	registry  *prometheus.Registry
	pusher    *push.Pusher
}

func metricName(app string, name string) string {
	return prometheus.BuildFQName(app, "", name)
}

func New(app string, endpoint string, usr string, pw string) *PrometheusPusher {
	monitor := push.New(endpoint, app)

	if len(usr) != 0 && len(pw) != 0 {
		monitor = monitor.BasicAuth(usr, pw)
	}

	r := prometheus.NewRegistry()

	m := PrometheusPusher{
		pusher:    monitor,
		registry:  r,
		metrics:   make(map[string]*prometheus.GaugeVec),
		labels:    make(map[string]string),
		labelKeys: make([]string, 0),
	}

	return &m
}

func (m *PrometheusPusher) Push() error {
	return m.pusher.Push()
}

func (m *PrometheusPusher) Configure() error {
	m.pusher.Gatherer(m.registry)

	return nil
}

func (m *PrometheusPusher) AddLabel(key string, value string) {
	m.labels[key] = value
	m.labelKeys = append(m.labelKeys, key)
}

func (m *PrometheusPusher) Register(name string, desc string) {
	m.metrics[name] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metricName(m.app, name),
			Help: desc,
		}, m.labelKeys,
	)
	m.registry.MustRegister(m.metrics[name])
}

func (m *PrometheusPusher) Increment(name string) {
	m.metrics[name].With(m.labels).Inc()
}

func (m *PrometheusPusher) Add(name string, delta float64) {
	m.metrics[name].With(m.labels).Add(delta)
}

func (m *PrometheusPusher) Reset(name string) {
	m.metrics[name].Reset()
}

func (m *PrometheusPusher) SetToCurrentTime(name string) {
	m.metrics[name].With(m.labels).SetToCurrentTime()
}
