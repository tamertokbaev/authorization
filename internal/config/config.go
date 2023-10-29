package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"sync"
)

type (
	Service struct {
		AppName     string         `envconfig:"APP_NAME" required:"true"`
		Port        string         `envconfig:"PORT" default:"8080"`
		Domain      string         `envconfig:"DOMAIN" default:"localhost"`
		Environment AppEnvironment `envconfig:"ENVIRONMENT" default:"local"`
	}

	Config struct {
		Service *Service
	}
)

var (
	once   sync.Once
	config *Config
)

type AppEnvironment string

const (
	PRODUCTION  AppEnvironment = "prod"
	STAGE       AppEnvironment = "stage"
	DEVELOPMENT AppEnvironment = "dev"
	LOCAL       AppEnvironment = "local"
)

func (e AppEnvironment) IsProduction() bool {
	return e == PRODUCTION
}

func (e AppEnvironment) IsStage() bool {
	return e == STAGE
}

func (e AppEnvironment) IsDevelopment() bool {
	return e == DEVELOPMENT
}

func (e AppEnvironment) IsLocal() bool {
	return e == LOCAL
}

func (e AppEnvironment) String() string {
	return string(e)
}

func (e AppEnvironment) Validate() error {
	switch e {
	case LOCAL, DEVELOPMENT, STAGE, PRODUCTION:
		return nil
	default:
		return fmt.Errorf("unexpected ENVIRONMENT in .env: %s", e)
	}
}

// GetConfig Загружает конфиг из .env файла и возвращает объект конфигурации
// В случае, если не передать параметр `envfiles`, берется `.env` файл из корня проекта
func GetConfig(envfiles ...string) (*Config, error) {
	var err error
	once.Do(func() {
		_ = godotenv.Load(envfiles...)

		var c Config
		err = envconfig.Process("", &c)
		if err != nil {
			err = fmt.Errorf("error parse config from env variables: %w\n", err)
			return
		}

		if e := c.Service.Environment.Validate(); e != nil {
			err = fmt.Errorf("error parse config from env variables: %w\n", e)
			return
		}

		config = &c
	})

	return config, err
}
