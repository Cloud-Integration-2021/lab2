package config

import (
	"fmt"
	"lab2/controllers"
	"lab2/models"

	"github.com/caarlos0/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppConfig struct {
	Env            string `env:"ENV" envDefault:"dev"`
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:"postgres"`
	Name     string `env:"DB_NAME" envDefault:"postgres"`
}

//IsDev return true if application is on dev stack
func (c *AppConfig) IsDev() bool {
	return IsDev(c.Env)
}

//IsDev return true if application is on dev stack
func IsDev(env string) bool {
	return env == "dev" || env == "development"
}

// LoadCfg loads the config file.
func LoadCfg() (AppConfig, error) {

	cfg := AppConfig{}
	env.Parse(&cfg)
	env.Parse(&cfg.DatabaseConfig)

	return cfg, nil
}

func (app *AppConfig) ConnectDatabase() (database controllers.Database, err error) {
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", app.DatabaseConfig.Host, app.DatabaseConfig.Port, app.DatabaseConfig.User, app.DatabaseConfig.Password, app.DatabaseConfig.Name)
	gorm, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	gorm.AutoMigrate(&models.Movie{})
	database.DB = gorm

	return database, nil
}
