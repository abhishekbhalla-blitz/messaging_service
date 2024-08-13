package config

// DefaultConfig
var defaultConfig = Config{
	Server: Server{
		Port:        8080,
		ContextPath: "/messaging-service",
	},
	Log: LogConfig{
		Level: "debug",
	},
	Producer: Producer{
		Primary: PrimaryConfig{
			Name:     "primary-cluster",
			Enabled:  true,
			FastFail: true,
			ProducerConfig: ProducerConfig{
				KafkaConfig: KafkaConfig{
					Brokers:          []string{"localhost:9092"},
					HealthCheckTopic: "health",
				},
			},
		},
		Fallback: FallbackConfig{
			Name:     "fallback-cluster",
			Enabled:  false,
			FastFail: true,
			ProducerConfig: ProducerConfig{
				KafkaConfig: KafkaConfig{
					Brokers:          []string{"localhost:9092"},
					HealthCheckTopic: "health",
				},
			},
		},
	},
}
