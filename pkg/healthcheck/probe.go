package healthcheck

import "sync/atomic"

type Probe struct {
	name    string
	route   string
	enabled atomic.Bool
}

func NewProbe(name string, initial bool, route string) *Probe {
	p := &Probe{name: name, route: route}
	p.enabled.Store(initial)
	return p
}

func NewLiveness() *Probe {
	return NewProbe("liveness", false, "/health/liveness")
}

func NewReadiness() *Probe {
	return NewProbe("readiness", false, "/health/readiness")
}

func (p *Probe) Name() string {
	return p.name
}

func (p *Probe) Route() string {
	return p.route
}

func (p *Probe) Enable() {
	p.enabled.Store(true)
}

func (p *Probe) Disable() {
	p.enabled.Store(false)
}

func (p *Probe) IsEnabled() bool {
	return p.enabled.Load()
}
