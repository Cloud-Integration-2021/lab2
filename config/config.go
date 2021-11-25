package config

import (
	"github.com/caarlos0/env"
)

type AppConfig struct {
	Env            string `env:"ENV" envDefault:"dev"`
}

//IsDev return true if application is on dev stack
func (app *AppConfig) IsDev() bool {
	return IsDev(app.Env)
}

//IsDev return true if application is on dev stack
func IsDev(env string) bool {
	return env == "dev" || env == "development"
}

// LoadCfg loads the config file.
func LoadCfg() (AppConfig, error) {

	cfg := AppConfig{}
	err := env.Parse(&cfg)
	if err != nil {
		return AppConfig{}, err
	}

	return cfg, nil
}

