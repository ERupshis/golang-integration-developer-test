package config

import (
	"errors"
	"flag"

	"github.com/caarlos0/env"
	"github.com/erupshis/golang-integration-developer-test/internal/common/utils/configutils"
)

// Config player's storage service.
type Config struct {
	Host        string `json:"address"`
	PlayersHost string `json:"players_host"`
	DatabaseDSN string `json:"database_dsn"`
	JWT         string `json:"jwt"`
	HashKey     string `json:"hash_key"`
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
	flagHostAddress        = "addr"
	flagPlayersHostAddress = "p_addr"
	flagDatabaseDSN        = "rdsn"
	flagJWT                = "jwt"
	flagHashKey            = "hk"
)

// checkFlags checks flags of app's launch.
func checkFlags(config *Config) {
	flag.StringVar(&config.Host, flagHostAddress, ":18081", "grpc server host")
	flag.StringVar(&config.PlayersHost, flagPlayersHostAddress, "localhost:8080", "players storage host")
	flag.StringVar(&config.DatabaseDSN, flagDatabaseDSN, "postgres://postgres:postgres@localhost:5432/auth_db?sslmode=disable", "records database DSN")
	flag.StringVar(&config.JWT, flagJWT, "SECRET_KEY", "jwt token generation key")
	flag.StringVar(&config.HashKey, flagHashKey, "SECRET_KEY", "user passwords hasher key")

	flag.Parse()
}

// ENVIRONMENTS PARSING.
// envConfig struct of environments suitable for agent.
type envConfig struct {
	Host        string `env:"HOST"`
	PlayersHost string `env:"PLAYERS_HOST"`
	DatabaseDSN string `env:"DATABASE_DSN"`
	JWT         string `env:"JWT_KEY"`
	HashKey     string `env:"HASH_KEY"`
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
	errs = append(errs, configutils.SetEnvToParamIfNeed(&config.PlayersHost, envs.PlayersHost))
	errs = append(errs, configutils.SetEnvToParamIfNeed(&config.DatabaseDSN, envs.DatabaseDSN))
	errs = append(errs, configutils.SetEnvToParamIfNeed(&config.JWT, envs.JWT))
	errs = append(errs, configutils.SetEnvToParamIfNeed(&config.HashKey, envs.HashKey))

	resErr := errors.Join(errs...)
	if resErr != nil {
		return configutils.ErrCheckEnvsWrapper(resErr)
	}

	return nil
}
