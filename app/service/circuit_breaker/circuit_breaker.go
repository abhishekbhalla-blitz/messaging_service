package circuit_breaker

import (
	log "github.com/sirupsen/logrus"
	"github.com/sony/gobreaker"
	"shopdeck.com/messaging_service/config"
	"time"
)

type CircuitBreaker struct {
	Breaker *gobreaker.CircuitBreaker
}

func CircuitBreakerInit(configuration config.Config) CircuitBreaker {
	cb := CircuitBreaker{}
	if configuration.Producer.Fallback.Enabled {
		cb.Breaker = getCircuitBreaker()
	}
	return cb
}

func getCircuitBreaker() *gobreaker.CircuitBreaker {
	cbSettings := gobreaker.Settings{
		Name:        "PrimaryKafkaCircuitBreaker",
		MaxRequests: 1,
		Interval:    60 * time.Second,
		Timeout:     30 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 10
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("Circuit breaker state changed from %s to %s\n", from, to)
		},
	}
	return gobreaker.NewCircuitBreaker(cbSettings)
}
