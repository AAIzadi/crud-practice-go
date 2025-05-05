package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Postgres PostgresConfig `yaml:"postgres"`
}

type ServerConfig struct {
	Port    string `yaml:"port"`
	RunMode string `yaml:"run_mode"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var config *Config
	err := v.Unmarshal(config)
	if err != nil {
		log.Printf("unable to decode struct, %v\n", err)
	}
	return config, nil
}

func localConfig(fileName string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath("./config")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		var fileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &fileNotFoundError) {
			return nil, errors.New("file not found")
		}
	}

	return v, nil
}

func getConfigPath(env string) string {
	if env == "production" {
		return "config-production"
	}
	if env == "staging" {
		return "config-staging"
	}
	if env == "local" {
		return "config-local"
	}
	panic("no valid env")
}

func GetConfig(env string) (*Config, error) {
	configPath := getConfigPath(env)
	loader, err := localConfig(configPath, "yaml")
	if err != nil {
		log.Printf("unable to finc config, %v\n", err)
		errors.New("no config found")
	}
	return parseConfig(loader)
}
