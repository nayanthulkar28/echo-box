package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		App  App  `mapstructure:",squash"`
		Http Http `mapstructure:",squash"`
		PG   PG   `mapstructure:",squash"`
	}

	PG struct {
		URL     string `mapstructure:"PG_URL"`
		PoolMax int    `mapstructure:"PG_POOLMAX"`
	}

	App struct {
		Name    string `mapstructure:"APP_NAME"`
		Version string `mapstructure:"APP_VERSION"`
	}

	Http struct {
		Port string `mapstructure:"HTTP_PORT"`
	}
)

func NewConfig(path string) (*Config, error) {
	var cfg *Config
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
