package httprequest

import "github.com/sony/gobreaker"

type HealthCheckResponse struct {
	ServerHealthStatus   ServerHealthStatus         `json:"server"`
	PrimaryHealthStatus  PrimaryHealthStatus        `json:"primaryCluster"`
	FallbackHealthStatus FallbackHealthStatus       `json:"fallbackCluster"`
	CircuitBreakerState  CircuitBreakerHealthStatus `json:"circuitBreakerState"`
}

type ServerHealthStatus struct {
	Healthy bool `json:"healthy"`
}

type PrimaryHealthStatus struct {
	Enabled bool `json:"enabled"`
	Healthy bool `json:"healthy"`
}

type FallbackHealthStatus struct {
	Enabled bool `json:"enabled"`
	Healthy bool `json:"healthy"`
}

type CircuitBreakerHealthStatus struct {
	Enabled bool            `json:"enabled"`
	State   gobreaker.State `json:"state"`
}

// HealthCheckResponseBuilder
type HealthCheckResponseBuilder struct {
	serverHealthStatus   ServerHealthStatus
	primaryHealthStatus  PrimaryHealthStatus
	fallbackHealthStatus FallbackHealthStatus
	circuitBreakerState  CircuitBreakerHealthStatus
}

func NewHealthCheckResponseBuilder() *HealthCheckResponseBuilder {
	return &HealthCheckResponseBuilder{}
}

func (healthCheckResponseBuilder *HealthCheckResponseBuilder) SetServerHealthStatus(healthy bool) *HealthCheckResponseBuilder {
	healthCheckResponseBuilder.serverHealthStatus.Healthy = healthy
	return healthCheckResponseBuilder
}

func (healthCheckResponseBuilder *HealthCheckResponseBuilder) SetPrimaryHealthStatus(enabled bool, healthy bool) *HealthCheckResponseBuilder {
	healthCheckResponseBuilder.primaryHealthStatus.Enabled = enabled
	healthCheckResponseBuilder.primaryHealthStatus.Healthy = healthy
	return healthCheckResponseBuilder
}

func (healthCheckResponseBuilder *HealthCheckResponseBuilder) SetFallbackHealthStatus(enabled bool, healthy bool) *HealthCheckResponseBuilder {
	healthCheckResponseBuilder.fallbackHealthStatus.Enabled = enabled
	healthCheckResponseBuilder.fallbackHealthStatus.Healthy = healthy
	return healthCheckResponseBuilder
}

func (healthCheckResponseBuilder *HealthCheckResponseBuilder) SetCircuitBreakerState(circuitBreakerState gobreaker.State) *HealthCheckResponseBuilder {
	healthCheckResponseBuilder.circuitBreakerState.Enabled = true
	healthCheckResponseBuilder.circuitBreakerState.State = circuitBreakerState
	return healthCheckResponseBuilder
}

func (healthCheckResponseBuilder *HealthCheckResponseBuilder) Build() HealthCheckResponse {
	return HealthCheckResponse{
		ServerHealthStatus:   healthCheckResponseBuilder.serverHealthStatus,
		PrimaryHealthStatus:  healthCheckResponseBuilder.primaryHealthStatus,
		FallbackHealthStatus: healthCheckResponseBuilder.fallbackHealthStatus,
		CircuitBreakerState:  healthCheckResponseBuilder.circuitBreakerState,
	}
}
