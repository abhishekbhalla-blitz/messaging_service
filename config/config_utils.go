package config

import "strings"

func (kafkaConfig *KafkaConfig) GetBrokers() string {
	return strings.Join(kafkaConfig.Brokers, ",")
}

func GetConfiguration() Config {
	return configuration
}
