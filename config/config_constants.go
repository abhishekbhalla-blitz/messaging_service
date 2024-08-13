package config

import (
	"fmt"
	"strings"
)

type Environment string

const (
	DEFAULT    Environment = "default"
	LOCAL      Environment = "local"
	STAGING    Environment = "staging"
	PREPROD    Environment = "prepreod"
	PRODUCTION Environment = "prod"
)

func ParseEnvironment(lvl string) (Environment, error) {
	switch strings.ToUpper(lvl) {
	case string(DEFAULT):
		return DEFAULT, nil
	case string(LOCAL):
		return LOCAL, nil
	case string(STAGING):
		return STAGING, nil
	case string(PREPROD):
		return PREPROD, nil
	case string(PRODUCTION):
		return PRODUCTION, nil
	default:
		return "", fmt.Errorf("invalid environment: %s", lvl)
	}
}
