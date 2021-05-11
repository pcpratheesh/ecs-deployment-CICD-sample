package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Environment string

func (e Environment) String() string {
	return string(e)
}

const (
	Local         Environment = "local"
	Development   Environment = "dev"
	IOS           Environment = "iosdev"
	Preproduction Environment = "preprod"
	Production    Environment = "prod"
)

var validEnvironments = []Environment{Local, Development, IOS, Preproduction, Production}

// all fields read from the environment, and prefixed with IMPART_
type Configuration struct {
	Env   Environment `split_words:"true" default:"dev"`
	Debug bool        `split_words:"true" default:"false"`
	Port  int         `split_words:"true" default:"8080"`
}

func GetConfig() (*Configuration, error) {
	var cfg Configuration
	if err := envconfig.Process("env", &cfg); err != nil {
		return nil, err
	}
	isValidEnvironment := false
	for _, e := range validEnvironments {
		if e == cfg.Env {
			isValidEnvironment = true
		}
	}
	if !isValidEnvironment {
		return nil, fmt.Errorf("invalid environment specified %s", cfg.Env)
	}

	return &cfg, nil
}
