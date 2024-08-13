package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

// Configuration is default + overrides from `config.yml`
var configuration Config

func InitConfig() {
	log.Info("Initializing Config")

	env := getEnvironment()
	log.Print("Current environment: ", env)

	configFile := "./config/" + getConfigFile(env)
	configData, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading init file: %v", err)
	}

	err = yaml.Unmarshal(configData, &configuration)
	if err != nil {
		log.Fatalf("Error parsing init file: %v", err)
	}

	logCurrentConfiguration()
}

func getEnvironment() Environment {
	var environment Environment = DEFAULT
	environmentVal, isEnvironmentSet := os.LookupEnv("env")

	if !isEnvironmentSet {
		log.Infof("Environment not set. Using default configuration.")
	} else {
		environment, err := ParseEnvironment(environmentVal)
		if err != nil {
			log.Fatalf("Invalid environment: %s", environmentVal)
		}
		log.Infof("Using environment: %s", environment)
	}
	return environment
}

func getConfigFile(env Environment) string {
	switch env {
	case LOCAL:
		return "config.local.yml"
	case STAGING:
		return "config.staging.yml"
	case PREPROD:
		return "config.preprod.yml"
	case PRODUCTION:
		return "config.prod.yml"
	default:
		return "config.default.yml"
	}
}

func logCurrentConfiguration() {
	level, err := log.ParseLevel(configuration.Log.Level)
	if err != nil {
		log.Fatalf("Invalid log level: %s", configuration.Log.Level)
	}

	if level >= log.DebugLevel {
		d, _err := yaml.Marshal(configuration)
		if _err != nil {
			log.Fatalf("Error marshaling configuration: %v", _err)
		}
		log.Print("Configuration:\n", string(d))
	}
}
