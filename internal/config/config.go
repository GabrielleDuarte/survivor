package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	ProjectName string      `mapstructure:"project_name"`
	HTTPPort    string      `mapstructure:"port"`
	Mongo       MongoConfig `mapstructure:"mongo"`
}

type MongoConfig struct {
	URI          string `mapstructure:"uri"`
	DatabaseName string `mapstructure:"database_name"`
	JWTSecret    string `mapstructure:"jwt_secret"`
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigName("local")
	v.SetConfigType("yaml")
	v.AddConfigPath("../internal/config/")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("config: read file: %w", err)
	}
	fmt.Println(v.ConfigFileUsed())

	v.AutomaticEnv()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("config: unmarshal: %w", err)
	}

	return &cfg, nil
}
