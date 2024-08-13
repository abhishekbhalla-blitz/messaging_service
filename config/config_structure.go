package config

// Config defines configuration structure
type Config struct {
	Server   Server    `yaml:"server"`
	Log      LogConfig `yaml:"log"`
	Producer Producer  `yaml:"producer"`
}

type Server struct {
	Port        int    `yaml:"port"`
	ContextPath string `yaml:"contextPath"`
}

type LogConfig struct {
	Level string `yaml:"level"`
}

type Producer struct {
	Primary  PrimaryConfig  `yaml:"primary"`
	Fallback FallbackConfig `yaml:"fallback"`
}

type PrimaryConfig struct {
	Name           string         `yaml:"name"`
	Enabled        bool           `yaml:"enabled"`  // enable producer for this cluster
	FastFail       bool           `yaml:"fastFail"` // fail initialization if initial connection fails
	ProducerConfig ProducerConfig `yaml:"producerConfig"`
}

type FallbackConfig struct {
	Name           string         `yaml:"name"`
	Enabled        bool           `yaml:"enabled"`
	FastFail       bool           `yaml:"fastFail"` // fail initialization if initial connection fails
	ProducerConfig ProducerConfig `yaml:"producerConfig"`
}

type ProducerConfig struct {
	Type        ProducerType `yaml:"type"`
	KafkaConfig KafkaConfig  `yaml:"kafkaConfig"`
}

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
	//Topic   string   `yaml:"topic"`
	HealthCheckTopic string `yaml:"healthCheckTopic"`
}

// ProducerType is a custom type defined to limit accepted values
type ProducerType string

const (
	KAFKA ProducerType = "kafka"
	//GCP_PUBSUB              = "pubsub"
	//AWS_SQS                 = "sqs"
)
