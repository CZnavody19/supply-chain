package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server        ServerConfig
	DBConfig      DBConfig
	LoggingConfig LoggingConfig
}

func LoadConfig() *Config {
	return &Config{
		Server:        loadServerConfig(),
		DBConfig:      loadDbConfig(),
		LoggingConfig: loadLoggingConfig(),
	}
}

func configViper(configName string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(strings.ToUpper(configName))
	v.AutomaticEnv()
	v.SetConfigName(configName)
	v.SetConfigType("yaml")
	v.AddConfigPath("./configuration/")
	return v
}
