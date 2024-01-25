package config

import (
	"errors"
	"flag"

	"github.com/caarlos0/env"
	"github.com/erupshis/golang-integration-developer-test/internal/common/utils/configutils"
)

// Config player's storage service.
type Config struct {
	Host string `json:"address"`
}

// Parse main func to parse variables.
func Parse() (Config, error) {
	var config = Config{}
	checkFlags(&config)
	err := checkEnvironments(&config)
	return config, err
}

// FLAGS PARSING.
const (
	flagHostAddress = "addr"
)

// checkFlags checks flags of app's launch.
func checkFlags(config *Config) {
	flag.StringVar(&config.Host, flagHostAddress, ":8081", "grpc server host")

	flag.Parse()
}

// ENVIRONMENTS PARSING.
// envConfig struct of environments suitable for agent.
type envConfig struct {
	Host string `env:"HOST"`
}

// checkEnvironments checks environments suitable for agent.
func checkEnvironments(config *Config) error {
	var envs = envConfig{}
	err := env.Parse(&envs)
	if err != nil {
		return configutils.ErrCheckEnvsWrapper(err)
	}

	var errs []error
	errs = append(errs, configutils.SetEnvToParamIfNeed(&config.Host, envs.Host))

	resErr := errors.Join(errs...)
	if resErr != nil {
		return configutils.ErrCheckEnvsWrapper(resErr)
	}

	return nil
}
