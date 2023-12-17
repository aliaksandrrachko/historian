package health

import (
	"fmt"
	"runtime"
	"time"
)

type HealthReport struct {
	Status     Status
	Components []Component
}

type Health struct {
	Status Status
}

type Component struct {
	Status      Status
	Description string
}

type Status int8

const (
	UNKNOWN Status = iota
	UP
	DOWN
	OUT_OF_SERVICE
)

func (status Status) String() (name string) {
	switch status {
	case UNKNOWN:
		name = "UNKNOWN"
	case UP:
		name = "UP"
	case DOWN:
		name = "DOWN"
	case OUT_OF_SERVICE:
		name = "OUT_OF_SERVICE"
	}
	return name
}

type HealthCheck interface {
	Liveness() HealthReport
	Readiness() HealthReport
}

type healthCheck struct {
	LivenessIndicators  []HealthIndicator
	ReadinessIndicators []HealthIndicator
}

type HealthCheckBuilder interface {
	AddLiveness(healthIndicator HealthIndicator) HealthCheckBuilder
	AddReadiness(healthIndicator HealthIndicator) HealthCheckBuilder
	Build() HealthCheck
}

type healthCheckBuilder struct {
	LivenessIndicators  []HealthIndicator
	ReadinessIndicators []HealthIndicator
}

func NewDefaultHealthCheck() HealthCheck {
	builder := NewHealthCheckBuilder()
	builder.AddLiveness(NewAlwaysReadyCheck())
	builder.AddReadiness(NewGCMaxPauseCheck())
	builder.AddReadiness(NewGoroutineCountCheck())
	return builder.Build()
}

func NewHealthCheckBuilder() healthCheckBuilder {
	return healthCheckBuilder{
		LivenessIndicators:  make([]HealthIndicator, 0),
		ReadinessIndicators: make([]HealthIndicator, 0),
	}
}

func (hcb *healthCheckBuilder) AddLiveness(healthIndicator HealthIndicator) HealthCheckBuilder {
	hcb.LivenessIndicators = append(hcb.LivenessIndicators, healthIndicator)
	return hcb
}

func (hcb *healthCheckBuilder) AddReadiness(healthIndicator HealthIndicator) HealthCheckBuilder {
	hcb.ReadinessIndicators = append(hcb.ReadinessIndicators, healthIndicator)
	return hcb
}

func (hcb healthCheckBuilder) Build() HealthCheck {
	return healthCheck(hcb)
}

func (hc healthCheck) Liveness() HealthReport {
	return hc.createReport(hc.LivenessIndicators)
}

func (hc healthCheck) Readiness() HealthReport {
	return hc.createReport(hc.ReadinessIndicators)
}

func (hc healthCheck) createReport(indicators []HealthIndicator) HealthReport {
	generalStatus := UP
	components := make([]Component, 0)
	for _, indicator := range indicators {
		component := Component{Status: indicator.Status(), Description: indicator.String()}
		components = append(components, component)
		if component.Status != UP {
			generalStatus = component.Status
		}
	}
	return HealthReport{Status: generalStatus, Components: components}
}

type HealthIndicator interface {
	String() string
	Status() Status
}

type GoroutineCountCheck struct {
	threshold uint64
}

func NewGoroutineCountCheck() GoroutineCountCheck {
	// todo delete hard coded threshold
	return GoroutineCountCheck{threshold: 100}
}

func (gcc GoroutineCountCheck) String() string {
	return "GoRoutineCountCheck"
}

func (gcc GoroutineCountCheck) Status() Status {
	count := runtime.NumGoroutine()
	if count > int(gcc.threshold) {
		return OUT_OF_SERVICE
	}
	return UP
}

type AlwaysReadyCheck struct {
}

func NewAlwaysReadyCheck() AlwaysReadyCheck {
	return AlwaysReadyCheck{}
}

func (arc AlwaysReadyCheck) String() string {
	return "AlwaysReadyCheck"
}

func (arc AlwaysReadyCheck) Status() Status {
	return UP
}

type GCMaxPauseCheck struct {
	threshold time.Duration
}

func NewGCMaxPauseCheck() GCMaxPauseCheck {
	// todo delete hard coded threshold
	return GCMaxPauseCheck{threshold: time.Minute}
}

func (gmp GCMaxPauseCheck) String() string {
	return "GCMaxPauseCheck"
}

func (gmp GCMaxPauseCheck) Status() Status {
	thresholdNanoseconds := uint64(gmp.threshold.Nanoseconds())
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	for _, pause := range stats.PauseNs {
		if pause > thresholdNanoseconds {
			fmt.Printf("Recent GC cycle took %s > %s", time.Duration(pause), gmp.threshold)
			return OUT_OF_SERVICE
		}
	}
	return UP
}
